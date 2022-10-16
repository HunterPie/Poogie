package backups

import "context"

type IBackupRepository interface {
	FindAllByUserId(ctx context.Context, userId string) []BackupUploadModel
	Save(ctx context.Context, userId string, uploadModel BackupUploadModel) (BackupUploadModel, error)
	DeleteById(ctx context.Context, userId string, id string) bool
}
