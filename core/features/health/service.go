package health

import (
	"context"

	"github.com/Haato3o/poogie/core/persistence/database"
)

type HealthService struct {
	db database.IDatabase
}

func NewService(db database.IDatabase) *HealthService {
	return &HealthService{db}
}

func (s *HealthService) IsHealthy(ctx context.Context) (bool, error) {
	return s.db.IsHealthy(ctx)
}
