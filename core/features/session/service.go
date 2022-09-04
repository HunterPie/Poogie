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

	token := s.authService.Create(user.Id)

	return token, nil
}

func (s *SessionService) RefreshSession(ctx context.Context, session string) (string, error) {

}
