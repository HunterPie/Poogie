package mongodb

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(uri, database string, isTracingEnabled bool) (*mongo.Client, error) {
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

	return client, nil
}
