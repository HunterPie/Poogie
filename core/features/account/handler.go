package account

import (
	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/Haato3o/poogie/pkg/crypto"
	"github.com/Haato3o/poogie/pkg/server"
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
		repository:    server.Database.GetAccountRepository(),
		cryptoService: crypto.NewCryptoService(server.Config.CryptoKey, server.Config.CryptoSalt),
		hashService:   crypto.NewHashService(server.Config.HashSalt),
	}
	controller := AccountController{
		service: &service,
	}

	router.POST("/account", controller.CreateNewAccountHandler)

	userRouter := router.Group("/user")

	authMiddleware := middlewares.NewUserTransformMiddleware()

	userRouter.Use(authMiddleware.TokenToUserIdTransform)

	userRouter.GET("/my", controller.GetMyUserHandler)
	userRouter.GET("/:userId", controller.GetUserHandler)

	return nil
}

func New() server.IRegisterableService {
	return &AccountHandler{}
}
