package mongodb

import (
	"context"
	"time"

	"github.com/Haato3o/poogie/core/persistence/supporter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const SUPPORTER_COLLECTION_NAME = "supporters"
const SUPPORTER_TOKEN_NOT_ACTIVATED = "TOKEN_NOT_ACTIVATED"

type SupporterSchema struct {
	UserId    string    `bson:"user_id"`
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	IsActive  bool      `bson:"is_active"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func toSupporterSchema(model supporter.SupporterModel) SupporterSchema {
	return SupporterSchema{
		UserId:    SUPPORTER_TOKEN_NOT_ACTIVATED,
		Email:     model.Email,
		Token:     model.Token,
		IsActive:  model.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s SupporterSchema) toSupporterModel() supporter.SupporterModel {
	return supporter.SupporterModel{
		Email:    s.Email,
		Token:    s.Token,
		IsActive: s.IsActive,
	}
}

type SupporterMongoRepository struct {
	*mongo.Collection
}

// RenewBy implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) RenewBy(ctx context.Context, email string) supporter.SupporterModel {
	query := bson.M{
		"email": email,
	}

	update := bson.M{
		"$set": bson.M{
			"is_active":  true,
			"updated_at": time.Now(),
		},
	}

	var document SupporterSchema
	r.FindOneAndUpdate(ctx, query, update).Decode(&document)

	return document.toSupporterModel()
}

func NewSupporterRepository(db *mongo.Database) *SupporterMongoRepository {
	return &SupporterMongoRepository{db.Collection(SUPPORTER_COLLECTION_NAME)}
}

// ExistsSupporter implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) ExistsSupporter(ctx context.Context, email string) bool {
	query := bson.M{
		"email": email,
	}

	var document SupporterSchema
	err := r.FindOne(ctx, query).Decode(&document)

	return err != mongo.ErrNoDocuments
}

// ExistsToken implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) ExistsToken(ctx context.Context, token string) bool {
	query := bson.M{
		"token":     token,
		"is_active": true,
	}

	var document SupporterSchema
	err := r.FindOne(ctx, query).Decode(&document)

	return err != mongo.ErrNoDocuments
}

// FindBy implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) FindBy(ctx context.Context, email string) supporter.SupporterModel {
	query := bson.M{
		"email": email,
	}

	var document SupporterSchema
	r.FindOne(ctx, query).Decode(&document)

	return document.toSupporterModel()
}

// Insert implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) Insert(ctx context.Context, supporter supporter.SupporterModel) supporter.SupporterModel {

	r.InsertOne(ctx, toSupporterSchema(supporter))

	return supporter
}

// RevokeBy implements supporter.ISupporterRepository
func (r *SupporterMongoRepository) RevokeBy(ctx context.Context, email string) supporter.SupporterModel {
	query := bson.M{
		"email": email,
	}

	update := bson.M{
		"$set": bson.M{
			"is_active":  false,
			"updated_at": time.Now(),
		},
	}

	var document supporter.SupporterModel
	r.FindOneAndUpdate(ctx, query, update).Decode(&document)

	return document
}
