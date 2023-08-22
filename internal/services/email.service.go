package services

type EmailService interface {
	SendEmail(to, subject, tmplPath string, data interface{}) (bool, error)
}
