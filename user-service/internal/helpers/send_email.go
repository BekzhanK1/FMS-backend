package helpers

import (
	"bytes"
	"html/template"
	"log"
	"user-service/internal/config"

	"gopkg.in/gomail.v2"
)

type OTPData struct {
	OtpCode string
}

// Helper function to load and parse the HTML template
func parseTemplate(templatePath string, data OTPData) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func SendEmail(to string, subject string, otpData OTPData, templatePath string) {
	go func() {
		m := gomail.NewMessage()
		m.SetHeader("From", config.Envs.MailUsername)
		m.SetHeader("To", to)
		m.SetHeader("Subject", subject)

		body, err := parseTemplate(templatePath, otpData)
		if err != nil {
			log.Printf("could not parse email template: %v", err)
			return
		}

		m.SetBody("text/html", body)

		d := gomail.NewDialer(config.Envs.MailServer, int(config.Envs.MailPort), config.Envs.MailUsername, config.Envs.MailPassword)

		if err := d.DialAndSend(m); err != nil {
			log.Printf("could not send email: %v", err)
			return
		}

		log.Printf("Email sent successfully to %s", to)
	}()
}
