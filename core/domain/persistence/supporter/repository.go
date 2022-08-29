package supporter

import "context"

type ISupporterRepository interface {
	FindBy(ctx context.Context, email string) SupporterModel
	Exists(ctx context.Context, email string) bool
	RevokeBy(ctx context.Context, email string) bool
	Insert(ctx context.Context, supporter SupporterModel) SupporterModel
}
