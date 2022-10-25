package mongodb

import (
	"context"
	"time"

	"github.com/Haato3o/poogie/core/persistence/patches"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PATCHES_COLLECTION_NAME = "patches"

type PatchSchema struct {
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	Link        string    `bson:"link"`
	Banner      string    `bson:"banner"`
	CreatedAt   time.Time `bson:"created_at"`
}

func (s PatchSchema) ToPatch() patches.Patch {
	return patches.Patch{
		Title:       s.Title,
		Description: s.Description,
		Link:        s.Link,
		Banner:      s.Banner,
	}
}

type PatchMongoRepository struct {
	*mongo.Collection
}

// FindAll implements patches.IPatchRepository
func (r *PatchMongoRepository) FindAll(ctx context.Context) []patches.Patch {
	var models = make([]patches.Patch, 0)

	options := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, _ := r.Find(ctx, bson.D{}, options)

	var schemas []PatchSchema
	cursor.All(ctx, &schemas)

	for _, schema := range schemas {
		models = append(models, schema.ToPatch())
	}

	return models
}

func NewPatchRepository(db *mongo.Database) patches.IPatchRepository {
	return &PatchMongoRepository{
		Collection: db.Collection(PATCHES_COLLECTION_NAME),
	}
}
