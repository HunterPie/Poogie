package smtp

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/Haato3o/poogie/core/email"
	"github.com/Haato3o/poogie/pkg/log"
)

const (
	SMTP_HOST = "smtp.gmail.com"
	SMTP_PORT = "587"
)

type SMTPEmail struct {
	Title        string
	Recipients   []string
	TemplatePath string
	TemplateData interface{}
}

type SMTPService struct {
	email    string
	password string
	queue    chan SMTPEmail
}

func New(email, password string) email.IEmailService {
	service := &SMTPService{email, password, make(chan SMTPEmail, 1000)}

	go service.handleQueue()
	return service
}

func (s *SMTPService) enqueue(smtpEmail SMTPEmail) {
	s.queue <- smtpEmail

	log.Info("enqueued email for " + smtpEmail.Recipients[0])
}

func (s *SMTPService) Send(title string, recipients []string, templatePath string, templateData interface{}) (bool, error) {

	s.enqueue(SMTPEmail{
		Title:        title,
		Recipients:   recipients,
		TemplatePath: templatePath,
		TemplateData: templateData,
	})

	return true, nil
}

func (s *SMTPService) handleQueue() {
	for {
		select {
		case email := <-s.queue:
			_, err := s.sendEmail(email)

			if err != nil {
				log.Error("failed to send email to "+email.Recipients[0], err)
				continue
			}

			log.Info("email sent to " + email.Recipients[0])
		}
	}
}

func (s *SMTPService) sendEmail(smtpEmail SMTPEmail) (bool, error) {
	auth := smtp.PlainAuth("", s.email, s.password, SMTP_HOST)

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", smtpEmail.Title, mimeHeaders)))

	temp, err := template.ParseFiles(smtpEmail.TemplatePath)

	if err != nil {
		return false, err
	}

	temp.Execute(&body, smtpEmail.TemplateData)

	if err := smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, s.email, smtpEmail.Recipients, body.Bytes()); err != nil {
		return false, err
	}

	return true, nil
}
