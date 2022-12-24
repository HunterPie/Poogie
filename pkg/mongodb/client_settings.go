package mongodb

import (
	"context"
	"github.com/Haato3o/poogie/core/persistence/settings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const ClientSettingsCollectionName = "client_settings"

type ClientSettingsSchema struct {
	UserId    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Settings  string             `bson:"settings"`
}

func (s ClientSettingsSchema) toClientSettingsModel() settings.ClientSettingModel {
	return settings.ClientSettingModel{
		UserId:          s.UserId.String(),
		EncodedSettings: s.Settings,
	}
}

type ClientSettingsMongoRepository struct {
	*mongo.Collection
}

func (c ClientSettingsMongoRepository) UpdateBy(
	ctx context.Context,
	userId string,
	encodedSettings string,
) (settings.ClientSettingModel, error) {
	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return settings.ClientSettingModel{}, ErrInvalidId
	}

	query := bson.M{
		"user_id": id,
	}

	_, err = c.FindOne(ctx, query).DecodeBytes()

	if err == mongo.ErrNoDocuments {
		schema := ClientSettingsSchema{
			UserId:    id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Settings:  encodedSettings,
		}

		c.InsertOne(ctx, schema)

		return schema.toClientSettingsModel(), nil
	} else if err != nil {
		return settings.ClientSettingModel{}, err
	}

	update := bson.M{
		"$set": bson.M{
			"settings":   encodedSettings,
			"updated_at": time.Now(),
		},
	}

	c.FindOneAndUpdate(ctx, query, update)

	var schema ClientSettingsSchema
	err = c.FindOne(ctx, query).Decode(&schema)

	return schema.toClientSettingsModel(), nil
}

func (c ClientSettingsMongoRepository) DeleteBy(ctx context.Context, userId string) bool {
	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return false
	}

	query := bson.M{
		"user_id": id,
	}

	c.FindOneAndDelete(ctx, query)

	return true
}

func (c ClientSettingsMongoRepository) FindBy(ctx context.Context, userId string) (settings.ClientSettingModel, error) {
	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return settings.ClientSettingModel{}, ErrFailedToFindUser
	}

	query := bson.M{
		"user_id": id,
	}

	var schema ClientSettingsSchema
	err = c.FindOne(ctx, query).Decode(&schema)

	if err != nil {
		return settings.ClientSettingModel{}, ErrFailedToFindUser
	}

	return schema.toClientSettingsModel(), nil
}

func NewClientSettingsRepository(db *mongo.Database) settings.IClientSettingsRepository {
	return &ClientSettingsMongoRepository{
		Collection: db.Collection(ClientSettingsCollectionName),
	}
}
