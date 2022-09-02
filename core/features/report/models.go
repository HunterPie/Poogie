package report

type CrashReportRequest struct {
	Version    string `json:"version" binding:"required"`
	GameBuild  string `json:"game_build" binding:"required"`
	Exception  string `json:"exception" binding:"required"`
	StackTrace string `json:"stacktrace" binding:"required"`
}

type CrashReportResponse struct {
	Ok bool `json:"ok"`
}
