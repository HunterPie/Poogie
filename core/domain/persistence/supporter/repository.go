package supporter

import "context"

type ISupporterRepository interface {
	FindBy(ctx context.Context, email string) SupporterModel
	ExistsSupporter(ctx context.Context, email string) bool
	ExistsToken(ctx context.Context, token string) bool
	RevokeBy(ctx context.Context, email string) SupporterModel
	Insert(ctx context.Context, supporter SupporterModel) SupporterModel
}
