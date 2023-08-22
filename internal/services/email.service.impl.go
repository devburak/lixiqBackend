package services

import (
	"lixIQ/backend/internal/utils"
	"log"

	"gopkg.in/gomail.v2"
)

type EmailServiceImpl struct {
	config utils.AppConfig
}

type GomailEmailService struct {
	sender string
	dialer *gomail.Dialer
}

func NewGomailEmailService(config *utils.AppConfig) EmailService {
	dialer := gomail.NewDialer(config.EmailHost, config.EmailPort, config.EmailFrom, config.EmailPassword)
	return &GomailEmailService{
		sender: config.EmailFrom,
		dialer: dialer,
	}
}

// func NewEmailServiceImpl(config utils.AppConfig) EmailServiceImpl {
// 	return EmailServiceImpl{config}
// }

// func (es EmailServiceImpl) SendEmail(to, subject, tmplPath string, data interface{}) (bool, error) {
// 	host := es.config.EmailHost
// 	from := es.config.EmailFrom
// 	password := es.config.EmailPassword

// 	fmt.Print(host)
// 	fmt.Print(", " + from)
// 	client, err := smtp.Dial(host)
// 	if err != nil {
// 		log.Fatal("Error on client Dial:", err)
// 		return false, err
// 	}

// 	tlsConfig := &tls.Config{
// 		InsecureSkipVerify: false,            // Sunucu sertifikasını doğrula
// 		ServerName:         "smtp.gmail.com", // E-posta sunucusunun adı
// 	}

// 	err = client.StartTLS(tlsConfig)
// 	if err != nil {
// 		log.Fatal("Error on client StartTLS:", err)
// 		return false, err
// 	}

// 	auth := smtp.PlainAuth("", from, password, host)

// 	if tmplPath != "" {
// 		tmpl, err := template.ParseFiles("../template/" + tmplPath)
// 		if err != nil {
// 			log.Fatal("Error on mail parse template :", err)
// 			return false, err
// 		}

// 		emailContent := new(strings.Builder)
// 		err = tmpl.Execute(emailContent, data)
// 		if err != nil {
// 			log.Fatal("Error on mail content template :", err)
// 			return false, err
// 		}

// 		err = smtp.SendMail(host, auth, from, []string{to}, []byte(emailContent.String()))
// 		if err != nil {
// 			log.Fatal("Error on sending mail template :", err)
// 			return false, err
// 		}
// 	} else {
// 		msg := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, subject)
// 		err = smtp.SendMail(host, auth, from, []string{to}, []byte(msg))
// 		if err != nil {
// 			log.Fatal("Error on sending mail message :", err)
// 			return false, err
// 		}
// 	}

// 	return true, nil
// }

func (ges *GomailEmailService) SendEmail(to, subject, tmplPath string, data interface{}) (bool, error) {

	// Load HTML template
	htmlTemplate, err := utils.LoadTemplate(tmplPath, data)
	if err != nil {
		log.Fatal("Error loading HTML template:", err)
	}

	message := gomail.NewMessage()
	message.SetHeader("From", ges.sender)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", htmlTemplate)

	if err := ges.dialer.DialAndSend(message); err != nil {
		log.Fatal("Error sending email:", err)
		return false, err // İşlem başarısız olduğunu ve hatayı döndür
	}

	return true, nil
}
