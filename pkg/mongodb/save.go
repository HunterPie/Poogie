package mongodb

import (
	"context"
	"time"

	"github.com/Haato3o/poogie/core/persistence/save"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const SAVES_COLLECTION_NAME = "backups"

type BackupSchema struct {
	Id        string        `bson:"id"`
	Type      save.GameType `bson:"type"`
	CreatedAt time.Time     `bson:"created_at"`
}

type UserSaveSchema struct {
	UserId  string         `bson:"user_id"`
	Backups []BackupSchema `bson:"backups"`
}

func toBackupSchema(model save.BackupModel) BackupSchema {
	return BackupSchema{
		Id:        model.Id,
		Type:      model.Type,
		CreatedAt: model.CreatedAt,
	}
}

func toBackupModels(schemas []BackupSchema) []save.BackupModel {
	models := make([]save.BackupModel, 0)

	for _, schema := range schemas {
		models = append(models, save.BackupModel{
			Id:        schema.Id,
			Type:      schema.Type,
			CreatedAt: schema.CreatedAt,
		})
	}

	return models
}

func toSaveModel(schema UserSaveSchema) save.SaveModel {
	return save.SaveModel{
		UserId:  schema.UserId,
		Backups: toBackupModels(schema.Backups),
	}
}

type SaveBackupMongoRepository struct {
	*mongo.Collection
}

// DeleteById implements save.SaveBackupRepository
func (r *SaveBackupMongoRepository) DeleteById(ctx context.Context, userId string, id string) {
	query := bson.M{
		"user_id": userId,
	}

	update := bson.M{
		"$pull": bson.M{
			"id": id,
		},
	}

	r.FindOneAndUpdate(ctx, query, update)
}

// GetAllByUserId implements save.SaveBackupRepository
func (r *SaveBackupMongoRepository) GetAllByUserId(ctx context.Context, userId string) (save.SaveModel, error) {
	query := bson.M{
		"user_id": userId,
	}

	var schema UserSaveSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	if err != nil {
		return save.SaveModel{}, err
	}

	return toSaveModel(schema), nil
}

// InsertBackup implements save.SaveBackupRepository
func (r *SaveBackupMongoRepository) InsertBackup(ctx context.Context, userId string, model save.BackupModel) (save.BackupModel, error) {
	query := bson.M{
		"user_id": userId,
	}

	update := bson.M{
		"$push": bson.M{
			"backups": toBackupSchema(model),
		},
	}

	var schema UserSaveSchema
	err := r.FindOneAndUpdate(ctx, query, update).Decode(&schema)

	if err != nil {
		return save.BackupModel{}, err
	}

	return model, nil
}

func NewBackupRepository(db *mongo.Database) save.ISaveBackupRepository {
	return &SaveBackupMongoRepository{db.Collection(SAVES_COLLECTION_NAME)}
}
