package database

import (
	"context"

	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/Haato3o/poogie/core/persistence/notifications"
	"github.com/Haato3o/poogie/core/persistence/save"
	"github.com/Haato3o/poogie/core/persistence/supporter"
)

type IDatabase interface {
	IsHealthy(ctx context.Context) (bool, error)
	GetSupporterRepository() supporter.ISupporterRepository
	GetAccountRepository() account.IAccountRepository
	GetSessionRepository() account.IAccountSessionRepository
	GetNotificationsRepository() notifications.INotificationRepository
	GetSaveBackupRepository() save.ISaveBackupRepository
}
