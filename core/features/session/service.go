package session

import (
	"context"
	"errors"

	"github.com/Haato3o/poogie/core/auth"
	"github.com/Haato3o/poogie/core/crypto"
	"github.com/Haato3o/poogie/core/persistence/account"
)

var (
	ErrWrongCredentials = errors.New("invalid username or password")
	ErrInvalidToken     = errors.New("invalid session token")
)

type SessionService struct {
	accountRepository account.IAccountRepository
	sessionRepository account.IAccountSessionRepository
	authService       auth.IAuthService
	hashService       crypto.IHashService
}

func (s *SessionService) CreateSession(ctx context.Context, credentials LoginRequest) (string, error) {
	hashedPassword := s.hashService.Hash(credentials.Password)

	isLoginValid := s.accountRepository.AreCredentialsValid(ctx, credentials.Username, hashedPassword)

	if !isLoginValid {
		return "", ErrWrongCredentials
	}

	user, _ := s.accountRepository.GetByUsername(ctx, credentials.Username)

	token, err := s.authService.Create(user.Id)

	hashedToken := s.hashService.Hash(token)
	s.sessionRepository.CreateSession(ctx, hashedToken)

	return token, err
}

func (s *SessionService) RevokeSession(ctx context.Context, token string) {
	hashedToken := s.hashService.Hash(token)
	s.sessionRepository.RevokeSession(ctx, hashedToken)
}
