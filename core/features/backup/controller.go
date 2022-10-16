package backup

import (
	"github.com/Haato3o/poogie/core/features/common"
	"github.com/Haato3o/poogie/core/persistence/backups"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

const SAVE_SIZE_LIMIT int64 = 120_000_000

type BackupController struct {
	*BackupService
}

func (c *BackupController) DeleteBackupFileHandler(ctx *gin.Context) {
	userId := utils.ExtractUserId(ctx)
	backupId := ctx.Param("backupId")

	if backupId == "" {
		http.ElementNotFound(ctx)
		return
	}

	success := c.DeleteBackupFile(ctx, userId, backupId)

	if !success {
		http.ElementNotFound(ctx)
		return
	}

	http.Ok(ctx, BackupDeleteResponse{
		BackupId: backupId,
	})
}

func (c *BackupController) GetAllBackupsHandler(ctx *gin.Context) {
	userId := utils.ExtractUserId(ctx)

	response := c.FindAllBackupsForUser(ctx, userId)

	http.Ok(ctx, response)
}

func (c *BackupController) UploadBackupHandler(ctx *gin.Context) {
	gameId := ctx.Param("gameId")

	if gameId == "" || !backups.IsGameType(gameId) {
		http.BadRequest(ctx, common.ErrInvalidBackupUpload)
		return
	}
	userId := utils.ExtractUserId(ctx)

	if !c.CanUserUpload(ctx, userId) {

		http.TooManyRequests(ctx, common.ErrBackupRateLimit)
		return
	}

	file, headers, err := ctx.Request.FormFile("file")

	if headers.Size > SAVE_SIZE_LIMIT {
		http.TooLarge(ctx, common.ErrBackupSizeTooLarge)
		return
	}

	if err != nil {
		http.BadRequest(ctx, common.ErrInvalidBackupUpload)
		return
	}

	response, err := c.UploadBackupFile(ctx, BackupUploadRequest{
		UserId: userId,
		Stream: file,
		Size:   headers.Size,
		Game:   backups.GameType(gameId),
	})

	if err != nil {
		http.InternalServerError(ctx)
		return
	}

	http.Ok(ctx, response)
}

func (c *BackupController) DownloadBackupHandler(ctx *gin.Context) {
	backupId := ctx.Param("backupId")
	userId := utils.ExtractUserId(ctx)

	if backupId == "" {
		http.ElementNotFound(ctx)
		return
	}

	backup, err := c.DownloadBackupFile(ctx, userId, backupId)

	if err != nil {
		http.InternalServerError(ctx)
		return
	}

	ctx.DataFromReader(200, backup.Size, backup.Type, backup.Reader, nil)
}
