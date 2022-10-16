package account

import (
	"context"
	"errors"
)

var (
	ErrFailedToCreateAccount = errors.New("failed to create new account")
	ErrFailedToFindAccount   = errors.New("failed to find account")
)

type IAccountRepository interface {
	Create(ctx context.Context, model AccountModel) (AccountModel, error)
	UpdateSupporterStatus(ctx context.Context, userId string, isSupporter bool) (AccountModel, error)
	AreCredentialsValid(ctx context.Context, email, password string) bool
	GetByEmail(ctx context.Context, email string) (AccountModel, error)
	GetById(ctx context.Context, userId string) (AccountModel, error)
	IsEmailTaken(ctx context.Context, email string) bool
	IsUsernameTaken(ctx context.Context, username string) bool
	DeleteBy(ctx context.Context, userId string) AccountModel
	UpdatePassword(ctx context.Context, userId, password string) AccountModel
	UpdateAvatar(ctx context.Context, userId, avatar string) AccountModel
	VerifyAccount(ctx context.Context, userId string)
}

type IAccountSessionRepository interface {
	CreateSession(ctx context.Context, token string) (string, error)
	RevokeSession(ctx context.Context, token string) string
	IsSessionValid(ctx context.Context, token string) bool
}

type IAccountBadgesRepository interface {
	Create(ctx context.Context, userId, badgeId string)
	Delete(ctx context.Context, userId, badgeId string)
}

type IAccountHuntStatisticSummaryRepository interface {
	Create(ctx context.Context, userId, badgeId string)
	Delete(ctx context.Context, userId, badgeId string)
}

type IAccountVerificationRepository interface {
	Create(ctx context.Context, token string, account string)
	Find(ctx context.Context, token string) (string, error)
}

type IAccountResetRepository interface {
	Create(ctx context.Context, code string, email string)
	IsTokenValid(ctx context.Context, code string, email string) bool
	RevokeBy(ctx context.Context, code string, email string)
}
