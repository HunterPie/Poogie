package mongodb

import (
	"context"
	"time"

	"github.com/Haato3o/poogie/core/persistence/account"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const SESSIONS_COLLECTION_NAME = "sessions"

type SessionSchema struct {
	Token     string    `bson:"token"`
	CreatedAt time.Time `bson:"created_at"`
}

type SessionMongoRepository struct {
	*mongo.Collection
}

// IsSessionValid implements account.IAccountSessionRepository
func (r *SessionMongoRepository) IsSessionValid(ctx context.Context, token string) bool {
	query := bson.M{
		"token": token,
	}

	var schema SessionSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	if err != nil && err != mongo.ErrNoDocuments {
		return false
	}

	return err != mongo.ErrNoDocuments
}

// CreateSession implements account.IAccountSessionRepository
func (r *SessionMongoRepository) CreateSession(ctx context.Context, token string) (string, error) {
	schema := SessionSchema{
		Token:     token,
		CreatedAt: time.Now(),
	}

	_, err := r.InsertOne(ctx, schema)

	if err != nil {
		return "", err
	}

	return token, nil
}

// RevokeSession implements account.IAccountSessionRepository
func (r *SessionMongoRepository) RevokeSession(ctx context.Context, token string) string {
	query := bson.M{
		"token": token,
	}

	r.FindOneAndDelete(ctx, query)

	return token
}

func NewSessionRepository(db *mongo.Database) account.IAccountSessionRepository {
	return &SessionMongoRepository{db.Collection(SESSIONS_COLLECTION_NAME)}
}
