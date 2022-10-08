package backup

import (
	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/Haato3o/poogie/pkg/aws"
	"github.com/Haato3o/poogie/pkg/crypto"
	"github.com/Haato3o/poogie/pkg/jwt"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
)

type BackupHandler struct{}

// GetName implements server.IRegisterableService
func (*BackupHandler) GetName() string {
	return "BackupHandler"
}

// GetVersion implements server.IRegisterableService
func (*BackupHandler) GetVersion() int {
	return server.V1
}

// Load implements server.IRegisterableService
func (*BackupHandler) Load(router *gin.RouterGroup, server *server.Server) error {
	service := (&BackupService{
		repository:        server.Database.GetBackupsRepository(),
		accountRepository: server.Database.GetAccountRepository(),
		bucket:            aws.New(server.Config, "backups/", ".zip"),
		DeleteJobQueue:    make(chan DeleteQueueMessage, 1000),
	}).Initialize()
	controller := BackupController{
		BackupService: service,
	}

	authMiddleware := middlewares.NewUserTransformMiddleware(
		jwt.New(server.Config.JwtKey),
		server.Database.GetSessionRepository(),
		crypto.NewHashService(
			server.Config.HashSalt,
		),
	)

	router.Use(authMiddleware.TokenToUserIdTransform)

	router.POST("/backup/upload/:gameId", controller.UploadBackupHandler)
	router.GET("/backup/:backupId", controller.DownloadBackupHandler)
	router.GET("/backup", controller.GetAllBackupsHandler)
	router.DELETE("/backup/:backupId", controller.DeleteBackupFileHandler)

	return nil
}

func New() server.IRegisterableService {
	return &BackupHandler{}
}
