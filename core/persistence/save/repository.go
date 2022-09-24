package save

import "context"

type ISaveBackupRepository interface {
	InsertBackup(ctx context.Context, userId string, model BackupModel) (BackupModel, error)
	DeleteById(ctx context.Context, userId string, id string)
	GetAllByUserId(ctx context.Context, userId string) (SaveModel, error)
}
