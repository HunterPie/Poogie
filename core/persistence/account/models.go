package account

import (
	"time"

	"github.com/Haato3o/poogie/core/domain"
)

type AccountBadgesModel struct {
	Id        string
	CreatedAt time.Time
}

type HuntStatisticsSummaryModel struct {
	Id        string
	CreatedAt time.Time
}

type AccountModel struct {
	Id                         string
	Username                   string
	Password                   string
	Email                      string
	ClientId                   string
	Experience                 int64
	Rating                     int64
	AvatarUri                  string
	Badges                     []AccountBadgesModel
	HuntStatisticsSummaryModel []HuntStatisticsSummaryModel
	IsSupporter                bool
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	LastSessionAt              time.Time
	IsArchived                 bool
	IsActive                   bool
}

func (m AccountModel) GetBackupRateLimit() time.Duration {
	if m.IsSupporter {
		return domain.BACKUP_RATE_LIMIT_SUPPORTER
	}

	return domain.BACKUP_RATE_LIMIT
}
