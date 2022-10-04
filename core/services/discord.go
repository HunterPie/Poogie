package services

import (
	"bytes"
	"mime/multipart"
	"net/http"
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
	defer multi.Close()

	part, err := multi.CreateFormFile("files[0]", "data.txt")

	if err != nil {
		return
	}

	part.Write([]byte(data))

	req, err := http.NewRequest(http.MethodPost, s.endpoint, &buffer)

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", multi.FormDataContentType())
	client := http.DefaultClient

	client.Do(req)
	req.Response.Body.Close()
}
