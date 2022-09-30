package log

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const ENDPOINT = "https://log-api.newrelic.com/log/v1?Api-Key="
const BUFFER_SIZE = 50

var (
	NewRelicLogger *NewRelicHeadlessLogger
)

type NewRelicLogMessage struct {
	Message string        `json:"message"`
	Error   string        `json:"error,omitempty"`
	Context []*LogContext `json:"context,omitempty"`
	Level   string        `json:"level"`
}

type NewRelicHeadlessLogger struct {
	apiKey string
	buffer []NewRelicLogMessage
}

func NewLogger(apiKey string) {
	NewRelicLogger = &NewRelicHeadlessLogger{apiKey: apiKey}

	Info("initialized NewRelic headless client")
}

func (l *NewRelicHeadlessLogger) Info(message string, ctx []*LogContext) {
	l.send(NewRelicLogMessage{
		Message: message,
		Context: ctx,
		Level:   "info",
	})
}

func (l *NewRelicHeadlessLogger) Error(message string, err error, ctx []*LogContext) {
	l.send(NewRelicLogMessage{
		Message: err.Error(),
		Context: ctx,
		Level:   "error",
	})
}

func (l *NewRelicHeadlessLogger) send(message NewRelicLogMessage) {
	if l.apiKey == "" {
		return
	}

	l.buffer = append(l.buffer, message)

	if len(l.buffer) > BUFFER_SIZE {
		var buffer bytes.Buffer
		encoder := json.NewEncoder(&buffer)
		encoder.Encode(l.buffer)

		req, err := http.NewRequest(http.MethodPost, ENDPOINT+l.apiKey, &buffer)

		if err != nil {
			Error("failed to flush log buffer", err)
		}

		client := &http.Client{}

		go client.Do(req)

		l.buffer = make([]NewRelicLogMessage, 0)
	}
}
