package mongodb

import (
	"context"
	"time"

	"github.com/Haato3o/poogie/core/persistence/account"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const ACCOUNT_RESET_COLLECTION_NAME = "account_reset"

type ResetCodeSchema struct {
	Email     string    `bson:"email"`
	Code      string    `bson:"code"`
	CreatedAt time.Time `bson:"created_at"`
}

type AccountResetMongoRepository struct {
	*mongo.Collection
}

// Create implements account.IAccountResetRepository
func (r *AccountResetMongoRepository) Create(ctx context.Context, code string, email string) {
	schema := ResetCodeSchema{
		Email:     email,
		Code:      code,
		CreatedAt: time.Time{},
	}

	r.InsertOne(ctx, schema)
}

// IsTokenValid implements account.IAccountResetRepository
func (r *AccountResetMongoRepository) IsTokenValid(ctx context.Context, code string, email string) bool {
	query := bson.M{
		"code":  code,
		"email": email,
	}

	var schema ResetCodeSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	return err == nil
}

// RevokeBy implements account.IAccountResetRepository
func (r *AccountResetMongoRepository) RevokeBy(ctx context.Context, code string, email string) {
	query := bson.M{
		"code":  code,
		"email": email,
	}

	r.FindOneAndDelete(ctx, query)
}

func NewResetRepository(db *mongo.Database) account.IAccountResetRepository {
	return &AccountResetMongoRepository{
		Collection: db.Collection(ACCOUNT_RESET_COLLECTION_NAME),
	}
}
