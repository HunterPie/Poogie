package account

import (
	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/Haato3o/poogie/pkg/aws"
	"github.com/Haato3o/poogie/pkg/crypto"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/Haato3o/poogie/pkg/smtp"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct{}

// GetName implements server.IRegisterableService
func (*AccountHandler) GetName() string {
	return "AccountHandler"
}

// GetVersion implements server.IRegisterableService
func (*AccountHandler) GetVersion() int {
	return server.V1
}

// Load implements server.IRegisterableService
func (*AccountHandler) Load(router *gin.RouterGroup, server *server.Server) error {
	service := AccountService{
		repository:             server.Database.GetAccountRepository(),
		resetRepository:        server.Database.GetAccountResetRepository(),
		supporterRepository:    server.Database.GetSupporterRepository(),
		cryptoService:          crypto.NewCryptoService(server.Config.CryptoKey, server.Config.CryptoSalt),
		hashService:            crypto.NewHashService(server.Config.HashSalt),
		verificationRepository: server.Database.GetAccountVerificationRepository(),
		emailService:           smtp.New(server.Config.PoogieEmail, server.Config.PoogiePassword),
		avatarStorage:          aws.New(server.Config, "avatars/", ""),
		cryptoRandom:           crypto.NewCryptoRandomService(),
	}
	controller := AccountController{
		service: &service,
	}

	router.POST("/account", controller.CreateNewAccountHandler)
	router.GET("/account/verify/:token", controller.VerifyAccountHandler)
	router.POST("/account/password/reset", controller.RequestPasswordResetHandler)
	router.POST("/account/password", controller.ChangePasswordHandler)

	userRouter := router.Group("/user")

	userRouter.Use(middlewares.BlockUnauthenticatedRequest)

	userRouter.GET("/me", controller.GetMyUserHandler)
	userRouter.GET("/:userId", controller.GetUserHandler)
	userRouter.POST("/avatar/upload", controller.UploadAvatarHandler)
	return nil
}

func New() server.IRegisterableService {
	return &AccountHandler{}
}
