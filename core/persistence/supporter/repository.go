package supporter

import "context"

type ISupporterRepository interface {
	FindBy(ctx context.Context, email string) SupporterModel
	FindByAssociation(ctx context.Context, userId string) SupporterModel
	AssociateToUser(ctx context.Context, email string, userId string) SupporterModel
	ExistsSupporter(ctx context.Context, email string) bool
	ExistsToken(ctx context.Context, token string) bool
	RevokeBy(ctx context.Context, email string) SupporterModel
	RenewBy(ctx context.Context, email string) SupporterModel
	Insert(ctx context.Context, model SupporterModel) SupporterModel
}
