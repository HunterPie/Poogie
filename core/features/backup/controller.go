package backup

import (
	"github.com/Haato3o/poogie/core/features/common"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

type BackupController struct {
	*BackupService
}

func (c *BackupController) UploadBackupHandler(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")

	if err != nil {
		http.BadRequest(ctx, common.ErrInvalidBackupUpload)
		return
	}

	c.UploadBackupFile(ctx, file)

	ctx.JSON(200, gin.H{"test": true})
}

func (c *BackupController) DownloadBackupHandler(ctx *gin.Context) {
	backupId := ctx.Param("backupId")

	if backupId == "" {
		http.ElementNotFound(ctx)
		return
	}

	backup, err := c.DownloadBackupFile(ctx, backupId)

	if err != nil {
		http.InternalServerError(ctx)
		return
	}

	ctx.DataFromReader(200, backup.Size, backup.Type, backup.Reader, nil)
}
