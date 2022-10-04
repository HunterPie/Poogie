package log

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const ENDPOINT = "https://log-api.newrelic.com/log/v1?Api-Key="
const BUFFER_SIZE = 100

var (
	NewRelicLogger *NewRelicHeadlessLogger
)

type NewRelicLogMessage struct {
	App     string        `json:"app"`
	Message string        `json:"message"`
	Error   string        `json:"error,omitempty"`
	Context []*LogContext `json:"context,omitempty"`
	Level   string        `json:"level"`
}

type NewRelicHeadlessLogger struct {
	apiKey string
	queue  chan NewRelicLogMessage
}

func NewLogger(apiKey string) {
	NewRelicLogger = &NewRelicHeadlessLogger{apiKey: apiKey, queue: make(chan NewRelicLogMessage, BUFFER_SIZE+1)}
	go NewRelicLogger.queueListener()

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

	message.App = "poogie-api:prod"

	l.queue <- message
}

func (l *NewRelicHeadlessLogger) queueListener() {
	buffer := make([]NewRelicLogMessage, 0)

	for {
		select {
		case message := <-l.queue:
			buffer = append(buffer, message)

			if len(buffer) >= BUFFER_SIZE {
				var payloadBuffer bytes.Buffer
				encoder := json.NewEncoder(&payloadBuffer)
				encoder.Encode(buffer)

				req, err := http.NewRequest(http.MethodPost, ENDPOINT+l.apiKey, &payloadBuffer)

				if err != nil {
					Error("failed to flush log buffer", err)
				}

				http.DefaultClient.Do(req)

				buffer = make([]NewRelicLogMessage, 0)
			}
		}
	}
}
