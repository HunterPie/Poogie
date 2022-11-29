package report

type CrashReportContext struct {
	RamTotal       uint64 `json:"ram_total"`
	RamUsed        int64  `json:"ram_used"`
	WindowsVersion string `json:"windows_version"`
}

type CrashReportRequest struct {
	Version    string             `json:"version" binding:"required"`
	GameBuild  string             `json:"game_build" binding:"required"`
	Exception  string             `json:"exception" binding:"required"`
	StackTrace string             `json:"stacktrace" binding:"required"`
	Context    CrashReportContext `json:"context"`
}

type CrashReportResponse struct {
	Ok bool `json:"ok"`
}
