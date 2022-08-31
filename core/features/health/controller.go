package health

import (
	"net/http"

	"github.com/Haato3o/poogie/core/features/health/presentation"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	*HealthService
}

func NewController(s *HealthService) *HealthController {
	return &HealthController{s}
}

func (controller *HealthController) GetServiceHealth(ctx *gin.Context) {
	isDbHealthy, err := controller.IsHealthy(ctx)

	if !isDbHealthy {
		ctx.JSON(http.StatusInternalServerError, presentation.IsHealthyResponse{
			Healthy: false,
			Error:   err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, presentation.IsHealthyResponse{
		Healthy: true,
	})
}
