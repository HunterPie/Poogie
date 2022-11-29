package backup

import (
	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/Haato3o/poogie/pkg/aws"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

// GetName implements server.IRegisterableService
func (*Handler) GetName() string {
	return "BackupHandler"
}

// GetVersion implements server.IRegisterableService
func (*Handler) GetVersion() int {
	return server.V1
}

// Load implements server.IRegisterableService
func (*Handler) Load(router *gin.RouterGroup, server *server.Server) error {
	service := (&Service{
		repository:        server.Database.GetBackupsRepository(),
		accountRepository: server.Database.GetAccountRepository(),
		bucket:            aws.New(server.Config, "backups/", ".zip"),
		DeleteJobQueue:    make(chan DeleteQueueMessage, 1000),
	}).Initialize()
	controller := Controller{
		Service: service,
	}

	backupRouter := router.Group("/backup")

	backupRouter.Use(middlewares.BlockUnauthenticatedRequest)
	{
		backupRouter.GET("/upload", controller.CanUserUploadHandler)
		backupRouter.POST("/upload/:gameId", controller.UploadBackupHandler)
		backupRouter.GET("", controller.GetAllBackupsHandler)
		backupRouter.GET("/:backupId", controller.DownloadBackupHandler)
		backupRouter.DELETE("/:backupId", controller.DeleteBackupFileHandler)
	}

	return nil
}

func New() server.IRegisterableService {
	return &Handler{}
}
