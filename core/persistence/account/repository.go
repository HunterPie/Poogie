package account

import "context"

type IAccountRepository interface {
	Create(ctx context.Context, model AccountModel) AccountModel
	GetById(ctx context.Context, userId string) (AccountModel, error)
	IsEmailTaken(ctx context.Context, email string) bool
	DeleteBy(ctx context.Context, userId string) AccountModel
	UpdatePassword(ctx context.Context, userId, password string) AccountModel
	UpdateAvatar(ctx context.Context, userId, avatar string) AccountModel
}

type IAccountSessionRepository interface {
	CreateSession(ctx context.Context, userId string) string
	RevokeSession(ctx context.Context, userId string) string
	GetUserIdBy(ctx context.Context, token string) (string, error)
}

type IAccountBadgesRepository interface {
	Create(ctx context.Context, userId, badgeId string)
	Delete(ctx context.Context, userId, badgeId string)
}

type IAccountHuntStatisticSummaryRepository interface {
	Create(ctx context.Context, userId, badgeId string)
	Delete(ctx context.Context, userId, badgeId string)
}
