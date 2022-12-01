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

	ok, _ := utils.DeserializeBody(ctx, &crash)

	if !ok {
		http.BadRequest(ctx, "FAILED_TO_RECEIVE_BODY")
		return
	}

	clientId := utils.ExtractClientId(ctx)

	c.service.SendCrashReport(ctx, crash, clientId)

	http.Ok(ctx, CrashReportResponse{true})
}
