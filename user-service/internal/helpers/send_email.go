package helpers

import (
	"html/template"
	"log"
	"strings"
	"user-service/internal/config"

	"gopkg.in/gomail.v2"
)

type OTPData struct {
	OtpCode string
}

func SendEmail(to string, subject string, otpData OTPData, tmpl *template.Template) {
	go func() {
		m := gomail.NewMessage()
		m.SetHeader("From", config.Envs.MailUsername)
		m.SetHeader("To", to)
		m.SetHeader("Subject", subject)

		// Generate the email body using the provided template and data
		var bodyBuilder strings.Builder
		err := tmpl.Execute(&bodyBuilder, otpData)
		if err != nil {
			log.Printf("could not execute email template: %v", err)
			return
		}

		m.SetBody("text/html", bodyBuilder.String())

		d := gomail.NewDialer(config.Envs.MailServer, int(config.Envs.MailPort), config.Envs.MailUsername, config.Envs.MailPassword)

		if err := d.DialAndSend(m); err != nil {
			log.Printf("could not send email: %v", err)
			return
		}

		log.Printf("Email sent successfully to %s", to)
	}()
}

