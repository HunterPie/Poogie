package mongodb

import (
	"context"
	"time"

	"github.com/Haato3o/poogie/core/persistence/notifications"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const NOTIFICATIONS_COLLECTION_NAME = "notifications"

type NotificationSchema struct {
	Title           string                         `bson:"title"`
	Message         string                         `bson:"message"`
	Icon            string                         `bson:"icon"`
	Type            notifications.NotificationType `bson:"notification_type"`
	PrimaryAction   string                         `bson:"primary_action"`
	SecondaryAction string                         `bson:"secondary_action"`
	CreatedAt       time.Time                      `bson:"created_at"`
}

func toModel(schema NotificationSchema) notifications.NotificationModel {
	return notifications.NotificationModel{
		Title:           schema.Title,
		Message:         schema.Message,
		Icon:            schema.Icon,
		Type:            schema.Type,
		PrimaryAction:   schema.PrimaryAction,
		SecondaryAction: schema.SecondaryAction,
		CreatedAt:       schema.CreatedAt,
	}
}

func toModels(schemas []NotificationSchema) []notifications.NotificationModel {
	models := make([]notifications.NotificationModel, len(schemas))

	for _, schema := range schemas {
		models = append(models, toModel(schema))
	}

	return models
}

type MongoNotificationsRepository struct {
	*mongo.Collection
}

// FindAll implements notifications.INotificationRepository
func (r *MongoNotificationsRepository) FindAll(ctx context.Context) []notifications.NotificationModel {
	var results []NotificationSchema
	cursor, _ := r.Find(ctx, bson.D{}, options.Find())

	cursor.All(ctx, &results)

	return toModels(results)
}

func NewNotificationsRepository(db *mongo.Database) notifications.INotificationRepository {
	return &MongoNotificationsRepository{db.Collection(NOTIFICATIONS_COLLECTION_NAME)}
}
