package domain

import (
	"time"
)

const (
	MAX_BACKUPS                 = 2
	MAX_BACKUPS_SUPPORTER       = 5
	BACKUP_RATE_LIMIT           = 3 * 24 * time.Hour
	BACKUP_RATE_LIMIT_SUPPORTER = 24 * time.Hour
)
