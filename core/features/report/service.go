package report

import (
	"context"
	"fmt"
	"strings"

	"github.com/Haato3o/poogie/core/services"
)

const CRASH_REPORT_TEMPLATE = `
ClientId: %s
Version: %s
Game Build: %s
Exception: %s
StackTrace: %s
`

type ReportService struct {
	webhookService *services.DiscordWebhookService
}

func (s *ReportService) SendCrashReport(ctx context.Context, report CrashReportRequest, clientId string) {
	reportFormatted := fmt.Sprintf(CRASH_REPORT_TEMPLATE,
		clientId,
		report.Version,
		report.GameBuild,
		report.Exception,
		unescapeString(report.StackTrace),
	)

	s.webhookService.Send(reportFormatted)
}

func unescapeString(text string) string {
	text = strings.Replace(text, "\\n", "\n", -1)
	return text
}
