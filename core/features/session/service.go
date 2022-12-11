package session

import (
	"context"
	"errors"
	"strings"

	"github.com/Haato3o/poogie/core/auth"
	"github.com/Haato3o/poogie/core/crypto"
	"github.com/Haato3o/poogie/core/persistence/account"
)

var (
	ErrWrongCredentials  = errors.New("invalid username or password")
	ErrInvalidToken      = errors.New("invalid session token")
	ErrUnverifiedAccount = errors.New("account is not verified")
)

type SessionService struct {
	accountRepository account.IAccountRepository
	sessionRepository account.IAccountSessionRepository
	authService       auth.IAuthService
	hashService       crypto.IHashService
	cryptoService     crypto.ICryptographyService
}

func (s *SessionService) CreateSession(ctx context.Context, credentials LoginRequest) (string, error) {
	hashedPassword := s.hashService.Hash(credentials.Password)

	caseInsensitiveEmail := strings.ToLower(credentials.Email)
	encryptedEmail := s.cryptoService.Encrypt(caseInsensitiveEmail)

	isLoginValid := s.accountRepository.AreCredentialsValid(ctx, encryptedEmail, hashedPassword)

	if !isLoginValid {
		return "", ErrWrongCredentials
	}

	user, _ := s.accountRepository.GetByEmail(ctx, encryptedEmail)

	if !user.IsActive {
		return "", ErrUnverifiedAccount
	}

	token, err := s.authService.Create(user.Id)

	hashedToken := s.hashService.Hash(token)
	s.sessionRepository.CreateSession(ctx, hashedToken)

	return token, err
}

func (s *SessionService) RevokeSession(ctx context.Context, token string) {
	hashedToken := s.hashService.Hash(token)
	s.sessionRepository.RevokeSession(ctx, hashedToken)
}
