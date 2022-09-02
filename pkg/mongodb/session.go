package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const SESSIONS_COLLECTION_NAME = "sessions"

type SessionMongoRepository struct {
	*mongo.Collection
}

// CreateSession implements account.IAccountSessionRepository
func (*SessionMongoRepository) CreateSession(ctx context.Context, userId string) string {
	panic("unimplemented")
}

// GetUserIdBy implements account.IAccountSessionRepository
func (*SessionMongoRepository) GetUserIdBy(ctx context.Context, token string) (string, error) {
	panic("unimplemented")
}

// RevokeSession implements account.IAccountSessionRepository
func (*SessionMongoRepository) RevokeSession(ctx context.Context, userId string) string {
	panic("unimplemented")
}

func NewSessionRepository(db *mongo.Database) *SessionMongoRepository {
	return &SessionMongoRepository{db.Collection(SESSIONS_COLLECTION_NAME)}
}
