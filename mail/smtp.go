package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendRawEmail(to string, subject string, body string) error {
	smtpHost := os.Getenv("SMTP_ENDPOINT")
	smtpPort := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	from := os.Getenv("SMTP_FROM")

	// Connect to the SMTP server
	conn, err := smtp.Dial(smtpHost + ":" + smtpPort)
	if err != nil {
		return err
	}

	// Issue the STARTTLS command
	tlsConfig := &tls.Config{
		ServerName:         smtpHost,
		InsecureSkipVerify: true,
	}
	if err := conn.StartTLS(tlsConfig); err != nil {
		log.Fatalf("failed to start TLS: %v", err)
	}

	// Authenticate with the server
	auth := smtp.PlainAuth("", username, password, smtpHost)
	if err = conn.Auth(auth); err != nil {
		return err
	}
	// Set the sender and recipient
	if err = conn.Mail(from); err != nil {
		return err
	}
	if err = conn.Rcpt(to); err != nil {
		return err
	}

	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		from, to, subject, body)

	// Send the email data
	w, err := conn.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	// Close the connection
	err = conn.Quit()
	if err != nil {
		return err
	}
	return nil
}
