package supporter

import (
	"github.com/Haato3o/poogie/core/services"
	"github.com/Haato3o/poogie/pkg/crypto"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/Haato3o/poogie/pkg/smtp"
	"github.com/gin-gonic/gin"
)

type SupporterHandler struct{}

// GetName implements server.IRegisterableService
func (*SupporterHandler) GetName() string {
	return "SupporterHandler"
}

// GetVersion implements server.IRegisterableService
func (*SupporterHandler) GetVersion() int {
	return server.V1
}

// Load implements server.IRegisterableService
func (h *SupporterHandler) Load(router *gin.RouterGroup, server *server.Server) error {
	service := SupporterService{
		repository:        server.Database.GetSupporterRepository(),
		emailService:      smtp.New(server.Config.PoogieEmail, server.Config.PoogiePassword),
		tokenService:      services.NewTokenService(),
		accountRepository: server.Database.GetAccountRepository(),
		cryptoService: crypto.NewCryptoService(
			server.Config.CryptoKey,
			server.Config.CryptoSalt,
		),
	}
	controller := SupporterController{
		service:        &service,
		patreonService: services.NewPatreonService(server.Config.PatreonWebhookSecret),
	}

	router.GET("/supporter/verify", controller.HandleVerifySupporter)
	router.POST("/supporter/webhook", controller.HandleSupporterWebhook)

	return nil
}

func New() server.IRegisterableService {
	return &SupporterHandler{}
}
