package report

import (
	"context"
	"fmt"
	"github.com/Haato3o/poogie/core/tracing"
	"strings"

	"github.com/Haato3o/poogie/core/services"
)

const CRASH_REPORT_TEMPLATE = `
ClientId: %s
Version: %s
Game Build: %s
Exception: %s
Total Ram: %d
Ram used: %d
Windows version: %s
StackTrace: %s
`

type ReportService struct {
	webhookService *services.DiscordWebhookService
}

func (s *ReportService) SendCrashReport(ctx context.Context, report CrashReportRequest, clientId string) {
	txn := tracing.FromContext(ctx)

	txn.AddProperty("exception_code", report.Exception)

	// TODO: Remove this later
	if strings.Contains(report.StackTrace, "set_ShutdownMode") {
		return
	}

	reportFormatted := fmt.Sprintf(CRASH_REPORT_TEMPLATE,
		clientId,
		report.Version,
		report.GameBuild,
		report.Exception,
		report.Context.RamTotal,
		report.Context.RamUsed,
		report.Context.WindowsVersion,
		unescapeString(report.StackTrace),
	)

	s.webhookService.Send(reportFormatted)
}

func unescapeString(text string) string {
	text = strings.Replace(text, "\\n", "\n", -1)
	return text
}
