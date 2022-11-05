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
	ErrFailedToUploadBackup = errors.New("failed to upload backup file")
)

type UploadRequest struct {
	UserId      string
	Stream      io.Reader
	Size        int64
	Game        backups.GameType
	IsSupporter bool
}

type DeleteQueueMessage struct {
	BackupId string
	UserId   string
}

type Service struct {
	repository        backups.IBackupRepository
	accountRepository account.IAccountRepository
	bucket            bucket.IBucket
	DeleteJobQueue    chan DeleteQueueMessage
}

func (s *Service) Initialize() *Service {
	go s.listenToJobQueue()
	return s
}

func (s *Service) CanUserUpload(ctx context.Context, userId string) bool {
	userBackups := s.repository.FindAllByUserId(ctx, userId)
	user, _ := s.accountRepository.GetById(ctx, userId)

	return !(len(userBackups) > 0 && time.Since(userBackups[0].UploadedAt) < user.GetBackupRateLimit())
}

func (s *Service) UploadBackupFile(ctx context.Context, request UploadRequest) (BackupResponse, error) {
	backupId := utils.NewRandomString()

	_, err := s.bucket.UploadFromStream(ctx, buildUserStorageName(request.UserId, backupId), request.Stream)

	userBackups := s.repository.FindAllByUserId(ctx, request.UserId)

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

	if request.IsSupporter {
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

func (s *Service) DownloadBackupFile(
	ctx context.Context,
	userId string,
	backupId string,
) (bucket.StreamedFile, error) {
	return s.bucket.DownloadToStream(ctx, buildUserStorageName(userId, backupId))
}

func (s *Service) FindAllBackupsForUser(
	ctx context.Context,
	userId string,
	isSupporter bool,
) UserBackupDetailsResponse {
	maxCount := domain.MAX_BACKUPS

	if isSupporter {
		maxCount = domain.MAX_BACKUPS_SUPPORTER
	}

	backupsModels := s.repository.FindAllByUserId(ctx, userId)

	return UserBackupDetailsResponse{
		Count:    len(backupsModels),
		MaxCount: maxCount,
		Backups:  ToBackupResponses(backupsModels),
	}
}

func (s *Service) DeleteBackupFile(ctx context.Context, userId string, backupId string) bool {
	success := s.repository.DeleteById(ctx, userId, backupId)

	if !success {
		return false
	}

	s.bucket.Delete(ctx, buildUserStorageName(userId, backupId))

	return true
}

func (s *Service) listenToJobQueue() {
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
