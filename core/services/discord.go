package services

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"time"
)

type DiscordWebhookService struct {
	endpoint string
}

func NewDiscordWebhookService(endpoint string) *DiscordWebhookService {
	return &DiscordWebhookService{endpoint: endpoint}
}

func (s *DiscordWebhookService) Send(data string) {
	var buffer bytes.Buffer
	multi := multipart.NewWriter(&buffer)

	part, err := multi.CreateFormFile("files[0]", "data.txt")

	if err != nil {
		return
	}

	part.Write([]byte(data))

	multi.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, &buffer)

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", multi.FormDataContentType())
	client := &http.Client{}

	client.Do(req)
}
