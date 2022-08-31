package smtp

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/Haato3o/poogie/core/email"
)

const (
	SMTP_HOST = "smtp.gmail.com"
	SMTP_PORT = "587"
)

type SMTPService struct {
	email    string
	password string
}

func New(email, password string) email.IEmailService {
	return &SMTPService{email, password}
}

func (s *SMTPService) Send(title string, recipients []string, templatePath string, templateData interface{}) (bool, error) {

	auth := smtp.PlainAuth("", s.email, s.password, SMTP_HOST)

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", title, mimeHeaders)))

	temp, err := template.ParseFiles(templatePath)

	if err != nil {
		return false, err
	}

	temp.Execute(&body, templateData)

	if err := smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, s.email, recipients, body.Bytes()); err != nil {
		return false, err
	}

	return true, nil
}
