package mongodb

import (
	"context"
	"time"

	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/Haato3o/poogie/core/persistence/database"
	"github.com/Haato3o/poogie/core/persistence/notifications"
	"github.com/Haato3o/poogie/core/persistence/supporter"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDatabase struct {
	*mongo.Client
	*mongo.Database
}

// GetAccountRepository implements database.IDatabase
func (m *MongoDatabase) GetAccountRepository() account.IAccountRepository {
	return NewAccountRepository(m.Database)
}

// GetSessionRepository implements database.IDatabase
func (m *MongoDatabase) GetSessionRepository() account.IAccountSessionRepository {
	return NewSessionRepository(m.Database)
}

// GetSupporterRepository implements database.IDatabase
func (m *MongoDatabase) GetSupporterRepository() supporter.ISupporterRepository {
	return NewSupporterRepository(m.Database)
}

// GetNotificationsRepository implements database.IDatabase
func (m *MongoDatabase) GetNotificationsRepository() notifications.INotificationRepository {
	return NewNotificationsRepository(m.Database)
}

// IsHealthy implements database.IDatabase
func (m *MongoDatabase) IsHealthy(ctx context.Context) (bool, error) {
	context, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := m.Ping(context, readpref.Primary())

	if err != nil {
		return false, err
	}

	return true, nil
}

func New(uri, database string, isTracingEnabled bool) (database.IDatabase, error) {
	config := options.Client().ApplyURI(uri)

	if isTracingEnabled {
		config.SetMonitor(nrmongo.NewCommandMonitor(nil))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, config)

	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, errors.Wrap(err, "timeout: could not connect to database")
	}

	return &MongoDatabase{client, client.Database(database)}, nil
}
