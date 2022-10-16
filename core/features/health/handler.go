package health

import (
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
)

const SCOPE = "/health"

type HealthHandler struct{}

func New() server.IRegisterableService {
	return &HealthHandler{}
}

func (h *HealthHandler) GetName() string {
	return "HealthHandler"
}

func (h *HealthHandler) GetVersion() int {
	return server.NO_VERSION
}

func (h *HealthHandler) Load(r *gin.RouterGroup, s *server.Server) error {
	service := NewService(s.Database)
	controller := NewController(service)

	r.GET(SCOPE, controller.GetServiceHealth)

	return nil
}
