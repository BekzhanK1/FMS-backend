package utils

import (
	"log"
	"user-service/internal/config"

	"gopkg.in/gomail.v2"
)



func SendEmail(to string, subject string, body string) {
	go func() {
		m := gomail.NewMessage()
		m.SetHeader("From", config.Envs.MailUsername)
		m.SetHeader("To", to)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", body)

		d := gomail.NewDialer(config.Envs.MailServer, int(config.Envs.MailPort), config.Envs.MailUsername, config.Envs.MailPassword)

		if err := d.DialAndSend(m); err != nil {
			log.Printf("could not send email: %v", err)  // Log the error if sending fails
			return
		}
		
		log.Printf("Email sent successfully to %s", to)
	}()
}