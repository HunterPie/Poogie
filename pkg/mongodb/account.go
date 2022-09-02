package mongodb

import (
	"context"

	"github.com/Haato3o/poogie/core/persistence/account"
	"go.mongodb.org/mongo-driver/mongo"
)

const ACCOUNTS_COLLECTION_NAME = "accounts"

type AccountMongoRepository struct {
	*mongo.Collection
}

// Create implements account.IAccountRepository
func (*AccountMongoRepository) Create(ctx context.Context, account account.AccountModel) account.AccountModel {
	panic("unimplemented")
}

// DeleteBy implements account.IAccountRepository
func (*AccountMongoRepository) DeleteBy(ctx context.Context, userId string) account.AccountModel {
	panic("unimplemented")
}

// GetById implements account.IAccountRepository
func (*AccountMongoRepository) GetById(ctx context.Context, userId string) (account.AccountModel, error) {
	panic("unimplemented")
}

// IsEmailTaken implements account.IAccountRepository
func (*AccountMongoRepository) IsEmailTaken(ctx context.Context, email string) bool {
	panic("unimplemented")
}

// UpdateAvatar implements account.IAccountRepository
func (*AccountMongoRepository) UpdateAvatar(ctx context.Context, userId string, avatar string) account.AccountModel {
	panic("unimplemented")
}

// UpdatePassword implements account.IAccountRepository
func (*AccountMongoRepository) UpdatePassword(ctx context.Context, userId string, password string) account.AccountModel {
	panic("unimplemented")
}

func NewAccountRepository(db *mongo.Database) *AccountMongoRepository {
	return &AccountMongoRepository{db.Collection(ACCOUNTS_COLLECTION_NAME)}
}
