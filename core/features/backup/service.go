package backup

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/Haato3o/poogie/core/domain"
	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/Haato3o/poogie/core/persistence/backups"
	"github.com/Haato3o/poogie/core/persistence/bucket"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/log"
)

var (
	ErrFailedToUploadBackup   = errors.New("failed to upload backup file")
	ErrBackupRateLimitReached = errors.New("backup rate limit for this account has been reached")
)

type BackupUploadRequest struct {
	UserId string
	Stream io.Reader
	Size   int64
	Game   backups.GameType
}

type DeleteQueueMessage struct {
	BackupId string
	UserId   string
}

type BackupService struct {
	repository        backups.IBackupRepository
	accountRepository account.IAccountRepository
	bucket            bucket.IBucket
	DeleteJobQueue    chan DeleteQueueMessage
}

func (s *BackupService) Initialize() *BackupService {
	go s.listenToJobQueue()
	return s
}

func (s *BackupService) CanUserUpload(ctx context.Context, userId string) bool {
	userBackups := s.repository.FindAllByUserId(ctx, userId)
	user, _ := s.accountRepository.GetById(ctx, userId)

	return !(len(userBackups) > 0 && time.Since(userBackups[0].UploadedAt) < user.GetBackupRateLimit())
}

func (s *BackupService) UploadBackupFile(ctx context.Context, request BackupUploadRequest) (BackupResponse, error) {
	backupId := utils.NewRandomString()

	_, err := s.bucket.UploadFromStream(ctx, buildUserStorageName(request.UserId, backupId), request.Stream)

	userBackups := s.repository.FindAllByUserId(ctx, request.UserId)
	user, _ := s.accountRepository.GetById(ctx, request.UserId)

	if err != nil {
		return BackupResponse{}, ErrFailedToUploadBackup
	}

	model, err := s.repository.Save(ctx, request.UserId, backups.BackupUploadModel{
		Id:         backupId,
		Size:       request.Size,
		Game:       request.Game,
		UploadedAt: time.Now(),
	})

	if err != nil {
		return BackupResponse{}, ErrFailedToUploadBackup
	}

	// Delete existing backups if there are more than the maximum for that account

	maxBackups := domain.MAX_BACKUPS

	if user.IsSupporter {
		maxBackups = domain.MAX_BACKUPS_SUPPORTER
	}

	if len(userBackups) >= maxBackups {
		userBackups = userBackups[maxBackups-1:]

		for _, backup := range userBackups {
			s.DeleteJobQueue <- DeleteQueueMessage{
				BackupId: backup.Id,
				UserId:   request.UserId,
			}
			log.Info("queueing delete of backup: " + backup.Id + " for userId: " + request.UserId)
		}
	}

	return ToBackupResponse(model), nil
}

func (s *BackupService) DownloadBackupFile(ctx context.Context, userId string, backupId string) (bucket.StreamedFile, error) {
	return s.bucket.DownloadToStream(ctx, buildUserStorageName(userId, backupId))
}

func (s *BackupService) FindAllBackupsForUser(ctx context.Context, userId string) UserBackupDetailsResponse {
	user, _ := s.accountRepository.GetById(ctx, userId)

	maxCount := domain.MAX_BACKUPS

	if user.IsSupporter {
		maxCount = domain.MAX_BACKUPS_SUPPORTER
	}

	backups := s.repository.FindAllByUserId(ctx, userId)

	return UserBackupDetailsResponse{
		Count:    len(backups),
		MaxCount: maxCount,
		Backups:  ToBackupResponses(backups),
	}
}

func (s *BackupService) DeleteBackupFile(ctx context.Context, userId string, backupId string) bool {
	success := s.repository.DeleteById(ctx, userId, backupId)

	if !success {
		return false
	}

	s.bucket.Delete(ctx, buildUserStorageName(userId, backupId))

	return true
}

func (s *BackupService) listenToJobQueue() {
	for message := range s.DeleteJobQueue {
		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)

		s.repository.DeleteById(ctx, message.UserId, message.BackupId)
		s.bucket.Delete(ctx, buildUserStorageName(message.UserId, message.BackupId))

		cancel()

		log.Info("deleted backup: " + buildUserStorageName(message.UserId, message.BackupId))
	}
}

func buildUserStorageName(userId, file string) string {
	return fmt.Sprintf("%s/%s", userId, file)
}
