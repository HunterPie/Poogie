package report

type CrashReportContext struct {
	RamTotal       uint64 `json:"ram_total"`
	RamUsed        int64  `json:"ram_used"`
	WindowsVersion string `json:"windows_version"`
}

type CrashReportRequest struct {
	Version    string             `json:"version"`
	GameBuild  string             `json:"game_build"`
	Exception  string             `json:"exception"`
	StackTrace string             `json:"stacktrace"`
	IsUiError  bool               `json:"is_ui_error"`
	Context    CrashReportContext `json:"context"`
}

type CrashReportResponse struct {
	Ok bool `json:"ok"`
}
