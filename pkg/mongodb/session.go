package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const SESSIONS_COLLECTION_NAME = "sessions"

type SessionSchema struct {
	UserId    string    `bson:"user_id"`
	Token     string    `bson:"token"`
	CreatedAt time.Time `bson:"created_at"`
	ExpiresAt time.Time `bson:"expires_at"`
}

type SessionMongoRepository struct {
	*mongo.Collection
}

// CreateSession implements account.IAccountSessionRepository
func (r *SessionMongoRepository) CreateSession(ctx context.Context, userId string, token string) (string, error) {
	schema := SessionSchema{
		UserId:    userId,
		Token:     token,
		CreatedAt: time.Time{},
		ExpiresAt: time.Time{},
	}

	_, err := r.InsertOne(ctx, schema)

	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserIdBy implements account.IAccountSessionRepository
func (r *SessionMongoRepository) GetUserIdBy(ctx context.Context, token string) (string, error) {
	query := bson.M{
		"token": token,
	}

	var schema SessionSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	if err != nil {
		return "", err
	}

	return schema.Token, nil
}

// RevokeSession implements account.IAccountSessionRepository
func (r *SessionMongoRepository) RevokeSession(ctx context.Context, token string) string {
	query := bson.M{
		"token": token,
	}

	r.FindOneAndDelete(ctx, query)

	return token
}

func NewSessionRepository(db *mongo.Database) *SessionMongoRepository {
	return &SessionMongoRepository{db.Collection(SESSIONS_COLLECTION_NAME)}
}
