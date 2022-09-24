package backup

import (
	"github.com/Haato3o/poogie/core/crypto"
	"github.com/Haato3o/poogie/core/persistence/save"
)

type BackupService struct {
	repository    save.ISaveBackupRepository
	cryptoService crypto.ICryptographyService
}
