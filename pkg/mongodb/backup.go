package mongodb

import (
	"context"
	"time"

	"github.com/Haato3o/poogie/core/persistence/backups"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const BACKUPS_COLLECTION_NAME = "backups"

type BackupSchema struct {
	Id        string           `bson:"id"`
	UserId    string           `bson:"user_id"`
	Size      int64            `bson:"size"`
	Game      backups.GameType `bson:"game"`
	CreatedAt time.Time        `bson:"created_at"`
}

func (s BackupSchema) ToUploadModel() backups.BackupUploadModel {
	return backups.BackupUploadModel{
		Id:         s.Id,
		Size:       s.Size,
		Game:       s.Game,
		UploadedAt: s.CreatedAt,
	}
}

func toUploadModels(schemas []BackupSchema) []backups.BackupUploadModel {
	var models = make([]backups.BackupUploadModel, 0)

	for _, schema := range schemas {
		models = append(models, schema.ToUploadModel())
	}

	return models
}

func toUploadSchema(userId string, model backups.BackupUploadModel) BackupSchema {
	return BackupSchema{
		Id:        model.Id,
		UserId:    userId,
		Size:      model.Size,
		Game:      model.Game,
		CreatedAt: model.UploadedAt,
	}
}

type BackupsMongoRepository struct {
	*mongo.Collection
}

// DeleteById implements backups.IBackupRepository
func (r *BackupsMongoRepository) DeleteById(ctx context.Context, userId string, id string) bool {
	query := bson.M{
		"id":      id,
		"user_id": userId,
	}

	var schema BackupSchema
	err := r.FindOneAndDelete(ctx, query).Decode(&schema)

	return err == nil
}

// FindAllByUserId implements backups.IBackupRepository
func (r *BackupsMongoRepository) FindAllByUserId(ctx context.Context, userId string) []backups.BackupUploadModel {
	query := bson.M{
		"user_id": userId,
	}
	options := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	var results []BackupSchema
	cursor, _ := r.Find(ctx, query, options)

	cursor.All(ctx, &results)

	return toUploadModels(results)
}

// Save implements backups.IBackupRepository
func (r *BackupsMongoRepository) Save(ctx context.Context, userId string, uploadModel backups.BackupUploadModel) (backups.BackupUploadModel, error) {
	schema := toUploadSchema(userId, uploadModel)

	r.InsertOne(ctx, schema)

	var insertedSchema BackupSchema
	query := bson.M{
		"id": uploadModel.Id,
	}
	err := r.FindOne(ctx, query).Decode(&insertedSchema)

	if err != nil {
		return backups.BackupUploadModel{}, err
	}

	return insertedSchema.ToUploadModel(), nil
}

func NewBackupsRepository(db *mongo.Database) backups.IBackupRepository {
	return &BackupsMongoRepository{
		Collection: db.Collection(BACKUPS_COLLECTION_NAME),
	}
}
