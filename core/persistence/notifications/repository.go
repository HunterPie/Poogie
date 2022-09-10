package notifications

import "context"

type INotificationRepository interface {
	FindAll(ctx context.Context) []NotificationModel
}
