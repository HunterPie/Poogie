package backup

import (
	"context"
	"io"

	"github.com/Haato3o/poogie/core/persistence/bucket"
)

type BackupService struct {
	bucket bucket.IBucket
}

func (s *BackupService) UploadBackupFile(ctx context.Context, stream io.Reader) bool {
	s.bucket.UploadFromStream(ctx, "testing", stream)
	return true
}

func (s *BackupService) DownloadBackupFile(ctx context.Context, backupId string) (bucket.StreamedFile, error) {
	return s.bucket.DownloadToStream(ctx, backupId)
}
