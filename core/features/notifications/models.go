package notifications

import (
	"time"

	"github.com/Haato3o/poogie/core/persistence/notifications"
)

type Notification struct {
	Title           string                         `json:"title"`
	Message         string                         `json:"message"`
	Icon            string                         `json:"icon"`
	Type            notifications.NotificationType `json:"notification_type"`
	PrimaryAction   string                         `json:"primary_action,omitempty"`
	SecondaryAction string                         `json:"secondary_action,omitempty"`
	CreatedAt       time.Time                      `json:"created_at"`
}

func fromModel(model notifications.NotificationModel) Notification {
	return Notification{
		Title:           model.Title,
		Message:         model.Message,
		Icon:            model.Icon,
		Type:            model.Type,
		PrimaryAction:   model.PrimaryAction,
		SecondaryAction: model.SecondaryAction,
		CreatedAt:       model.CreatedAt,
	}
}
