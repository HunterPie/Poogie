package account

import (
	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/Haato3o/poogie/pkg/aws"
	"github.com/Haato3o/poogie/pkg/crypto"
	"github.com/Haato3o/poogie/pkg/jwt"
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
		supporterRepository:    server.Database.GetSupporterRepository(),
		cryptoService:          crypto.NewCryptoService(server.Config.CryptoKey, server.Config.CryptoSalt),
		hashService:            crypto.NewHashService(server.Config.HashSalt),
		verificationRepository: server.Database.GetAccountVerificationRepository(),
		emailService: smtp.New(
			server.Config.PoogieEmail,
			server.Config.PoogiePassword,
		),
		avatarStorage: aws.New(server.Config, "avatars/", ""),
	}
	controller := AccountController{
		service: &service,
	}

	router.POST("/account", controller.CreateNewAccountHandler)
	router.GET("/account/verify/:token", controller.VerifyAccount)

	userRouter := router.Group("/user")

	authMiddleware := middlewares.NewUserTransformMiddleware(
		jwt.New(server.Config.JwtKey),
		server.Database.GetSessionRepository(),
		crypto.NewHashService(
			server.Config.HashSalt,
		),
	)

	userRouter.Use(authMiddleware.TokenToUserIdTransform)

	userRouter.GET("/me", controller.GetMyUserHandler)
	userRouter.GET("/:userId", controller.GetUserHandler)
	userRouter.POST("/avatar/upload", controller.UploadAvatar)
	return nil
}

func New() server.IRegisterableService {
	return &AccountHandler{}
}
