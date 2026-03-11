package mailservice

import (
	"os"
	logger "smtpconnect/internal/Helper/Logger"
	"strconv"

	"gopkg.in/gomail.v2"
)

func MailService(toMailer string, htmlContent string, subject string, ccMailers []string) bool {
	log := logger.InitLogger()
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAILID"))
	m.SetHeader("To", toMailer)

	// ✅ Set CC only if provided
	if len(ccMailers) > 0 {
		m.SetHeader("Cc", ccMailers...)
	}

	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlContent)

	port, err := strconv.Atoi(os.Getenv("MAILPORT"))
	if err != nil {
		port = 587
	}

	// Use Gmail SMTP server
	d := gomail.NewDialer(os.Getenv("MAILCONNECTION"), port, os.Getenv("EMAILID"), os.Getenv("EMAILPASSWORD"))

	// Use TLS explicitly (optional, 587 already uses STARTTLS)
	d.TLSConfig = nil

	if err := d.DialAndSend(m); err != nil {
		log.Errorf("Could not send email: %v", err)
		return false
	}

	log.Println("Email sent successfully!")
	return true
}
