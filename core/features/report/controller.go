package report

import (
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
	service *ReportService
}

func (c *ReportController) SendCrashReportHandler(ctx *gin.Context) {
	var crash CrashReportRequest

	utils.DeserializeBody(ctx, &crash)
	clientId := utils.ExtractClientId(ctx)

	c.service.SendCrashReport(ctx, crash, clientId)

	http.Ok(ctx, CrashReportResponse{true})
}
