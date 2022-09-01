package supporter

import (
	"github.com/Haato3o/poogie/core/services"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/Haato3o/poogie/pkg/smtp"
	"github.com/gin-gonic/gin"
)

const SCOPE = "/supporter"

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
		repository:   server.Database.GetSupporterRepository(),
		emailService: smtp.New(server.Config.PoogieEmail, server.Config.PoogiePassword),
		tokenService: services.NewTokenService(),
	}
	controller := SupporterController{
		service:        &service,
		patreonService: services.NewPatreonService(server.Config.PatreonWebhookSecret),
	}

	router.GET(subScope("/verify"), controller.VerifySupporter)
	router.POST(subScope("/webhook"), controller.HandleSupporterWebhook)

	return nil
}

func New() server.IRegisterableService {
	return &SupporterHandler{}
}

func subScope(path string) string {
	return SCOPE + path
}
