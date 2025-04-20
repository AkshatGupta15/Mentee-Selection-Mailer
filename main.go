package main

import (
	"Personalised_Mailer/utils"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func main() {
	cfg := utils.LoadConfig()
	recipients, err := utils.ReadEmailsFromCSV("sample.csv")
	if err != nil {
		log.Fatalf("Failed to read emails from CSV: %v", err)
	}
	subject := fmt.Sprintf("["+cfg.Club.Name+"] Summer Project Selection Results - %s", cfg.Project.Name)

	auth := smtp.PlainAuth("", cfg.Mail.User, cfg.Mail.Pass, cfg.Mail.Host)

	for _, recipient := range recipients {
		body := utils.BuildEmailHTML(cfg, recipient.Name)

		var message strings.Builder
		message.WriteString(fmt.Sprintf("From: %s\r\n", cfg.Mail.User))
		message.WriteString(fmt.Sprintf("To: %s\r\n", recipient.Email+"@iitk.ac.in"))
		message.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
		message.WriteString("MIME-Version: 1.0\r\n")
		message.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n")
		message.WriteString(body)

		// Send the email
		err := smtp.SendMail(cfg.Mail.Host+":"+cfg.Mail.Port, auth, cfg.Mail.User, []string{recipient.Email}, []byte(message.String()))
		if err != nil {
			log.Printf("Failed to send mail to %s: %v", recipient.Email, err)
		} else {
			log.Printf("Email sent successfully to %s", recipient.Email)
		}
	}
}
