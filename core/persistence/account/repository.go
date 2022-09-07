package account

import (
	"context"
	"errors"
)

var (
	ErrFailedToCreateAccount = errors.New("failed to create new account")
)

type IAccountRepository interface {
	Create(ctx context.Context, model AccountModel) (AccountModel, error)
	AreCredentialsValid(ctx context.Context, username, password string) bool
	GetByUsername(ctx context.Context, username string) (AccountModel, error)
	GetById(ctx context.Context, userId string) (AccountModel, error)
	IsEmailTaken(ctx context.Context, email string) bool
	IsUsernameTaken(ctx context.Context, username string) bool
	DeleteBy(ctx context.Context, userId string) AccountModel
	UpdatePassword(ctx context.Context, userId, password string) AccountModel
	UpdateAvatar(ctx context.Context, userId, avatar string) AccountModel
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
