package notifications

import "time"

const (
	WARNING = "WARNING"
	INFO    = "INFO"
)

type NotificationType string

type NotificationModel struct {
	Title           string           `json:"title" bson:"title"`
	Message         string           `json:"message" bson:"message"`
	Icon            string           `json:"icon" bson:"icon"`
	Type            NotificationType `json:"notification_type" bson:"notification_type"`
	PrimaryAction   string           `json:"primary_action,omitempty" bson:"primary_action"`
	SecondaryAction string           `json:"secondary_action,omitempty" bson:"secondary_action"`
	CreatedAt       time.Time        `json:"created_at" bson:"created_at"`
}
