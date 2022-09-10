package notifications

import (
	"context"

	"github.com/Haato3o/poogie/core/persistence/notifications"
)

type NotificationsService struct {
	repository notifications.INotificationRepository
}

func (s *NotificationsService) GetAllNotifications(ctx context.Context) []Notification {
	models := s.repository.FindAll(ctx)
	notifications := make([]Notification, len(models))

	for _, model := range models {
		notifications = append(notifications, fromModel(model))
	}

	return notifications
}
