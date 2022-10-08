package backup

import (
	"github.com/Haato3o/poogie/core/features/common"
	"github.com/Haato3o/poogie/core/persistence/backups"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

const SAVE_SIZE_LIMIT int64 = 53_000_000

type BackupController struct {
	*BackupService
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

	file, headers, err := ctx.Request.FormFile("file")

	if headers.Size > SAVE_SIZE_LIMIT {
		http.TooLarge(ctx, common.ErrBackupSizeTooLarge)
		return
	}

	userId := utils.ExtractUserId(ctx)

	if err != nil {
		http.BadRequest(ctx, common.ErrInvalidBackupUpload)
		return
	}

	response, _ := c.UploadBackupFile(ctx, BackupUploadRequest{
		UserId: userId,
		Stream: file,
		Size:   headers.Size,
		Game:   backups.GameType(gameId),
	})

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
