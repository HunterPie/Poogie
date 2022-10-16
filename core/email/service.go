package email

type IEmailService interface {
	Send(title string, recipients []string, templatePath string, templateData interface{}) (bool, error)
}
