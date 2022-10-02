package mongodb

import (
	"context"
	"time"

	"github.com/Haato3o/poogie/core/persistence/account"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const ACCOUNT_VERIFICATION_COLLECTION = "account_verifications"

type AccountVerificationMongoRepository struct {
	*mongo.Collection
}

type AccountVerificationSchema struct {
	Token     string    `bson:"token"`
	AccountId string    `bson:"account_id"`
	CreatedAt time.Time `bson:"created_at"`
}

// Create implements account.IAccountVerificationRepository
func (r *AccountVerificationMongoRepository) Create(ctx context.Context, token string, account string) {
	r.InsertOne(ctx, AccountVerificationSchema{
		Token:     token,
		AccountId: account,
		CreatedAt: time.Now(),
	})
}

// Find implements account.IAccountVerificationRepository
func (r *AccountVerificationMongoRepository) Find(ctx context.Context, token string) (string, error) {

	query := bson.M{
		"token": token,
	}

	var schema AccountVerificationSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	if err != nil {
		return "", err
	}

	return schema.AccountId, nil
}

func NewAccountVerificationRepository(db *mongo.Database) account.IAccountVerificationRepository {
	return &AccountVerificationMongoRepository{db.Collection(ACCOUNT_VERIFICATION_COLLECTION)}
}
