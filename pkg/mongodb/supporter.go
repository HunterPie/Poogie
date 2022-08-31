package mongodb

import (
	"context"

	"github.com/Haato3o/poogie/core/persistence/supporter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const SUPPORTER_COLLECTION_NAME = "supporters"

type SupporterMongoRepository struct {
	*mongo.Collection
}

func NewSupporterRepository(db *mongo.Database) *SupporterMongoRepository {
	return &SupporterMongoRepository{db.Collection(SUPPORTER_COLLECTION_NAME)}
}

// ExistsSupporter implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) ExistsSupporter(ctx context.Context, email string) bool {
	query := bson.M{
		"email": email,
	}

	var document supporter.SupporterModel
	err := r.FindOne(ctx, query).Decode(&document)

	return err != mongo.ErrNoDocuments
}

// ExistsToken implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) ExistsToken(ctx context.Context, token string) bool {
	query := bson.M{
		"token": token,
	}

	var document supporter.SupporterModel
	err := r.FindOne(ctx, query).Decode(&document)

	return err != mongo.ErrNoDocuments
}

// FindBy implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) FindBy(ctx context.Context, email string) supporter.SupporterModel {
	query := bson.M{
		"email": email,
	}

	var document supporter.SupporterModel
	r.FindOne(ctx, query).Decode(&document)

	return document
}

// Insert implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) Insert(ctx context.Context, supporter supporter.SupporterModel) supporter.SupporterModel {

	r.InsertOne(ctx, supporter)

	return supporter
}

// RevokeBy implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) RevokeBy(ctx context.Context, email string) supporter.SupporterModel {
	query := bson.M{
		"email": email,
	}

	update := bson.M{
		"$set": bson.M{
			"is_active": false,
		},
	}

	var document supporter.SupporterModel
	r.FindOneAndUpdate(ctx, query, update).Decode(&document)

	return document
}
