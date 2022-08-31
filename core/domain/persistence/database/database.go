package database

import (
	"context"

	"github.com/Haato3o/poogie/core/domain/persistence/supporter"
)

type IDatabase interface {
	IsHealthy(ctx context.Context) (bool, error)
	GetSupporterRepository() supporter.ISupporterRepository
}
