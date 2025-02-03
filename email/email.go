package email

import (
	"encoding/base64"
	"fmt"
	"io"

	"gopkg.in/gomail.v2"
)

// EmailAttachment represents an email attachment encoded in Base64.
//
// Fields:
// - Base64: The Base64-encoded content of the file.
// - Name: The name of the attachment file.
type EmailAttachment struct {
	Base64 string
	Name   string
}

// SendEmailProps contains the properties required to send an email.
//
// Fields:
// - To: The primary recipient of the email.
// - Subject: The subject of the email.
// - Body: The body of the email in HTML format.
// - Attachments: A list of attachments to include in the email.
// - Cc: A list of CC (Carbon Copy) recipients.
// - Bcc: A list of BCC (Blind Carbon Copy) recipients.
type SendEmailProps struct {
	To          string
	Subject     string
	Body        string
	Attachments []EmailAttachment
	Cc          []string
	Bcc         []string
}

// EmailService defines the interface for sending emails.
type EmailService interface {
	// SendEmail sends an email with the given properties.
	//
	// Parameters:
	// - props: The email properties, including recipients, subject, body, and attachments.
	//
	// Returns:
	// - An error if the email fails to send.
	SendEmail(props SendEmailProps) error
}

// EmailServiceImpl implements the EmailService interface.
type EmailServiceImpl struct {
	Dialer *gomail.Dialer // SMTP dialer for sending emails
}

// MakeEmailServiceImpl creates a new EmailServiceImpl instance.
//
// Parameters:
// - user: The SMTP username (usually your email address).
// - password: The SMTP password (or app-specific password for Gmail).
// - host: The SMTP server host (e.g., "smtp.gmail.com").
// - port: The SMTP server port (e.g., 587 for Gmail).
//
// Returns:
// - A pointer to a new EmailServiceImpl instance.
func MakeEmailServiceImpl(
	user string, password string, host string, port int,
) *EmailServiceImpl {
	dialer := gomail.NewDialer(host, port, user, password)

	return &EmailServiceImpl{
		Dialer: dialer,
	}
}

// SendEmail sends an email using the provided properties.
//
// Parameters:
// - props: The email properties, including recipients, subject, body, and attachments.
//
// Returns:
// - An error if any step fails (e.g., SMTP connection, attachment decoding, or email sending).
//
// Example usage:
//
//	props := SendEmailProps{
//	    To:      "recipient@example.com",
//	    Subject: "Test Email",
//	    Body:    "<h1>Hello from Go!</h1><p>This is a test email.</p>",
//	    Attachments: []EmailAttachment{
//	        {
//	            Base64: "VGhpcyBpcyBhIHRlc3QgZmlsZSBhdHRhY2htZW50Lg==",
//	            Name:   "test.txt",
//	        },
//	    },
//	    Cc:  []string{"cc@example.com"},
//	    Bcc: []string{"bcc@example.com"},
//	}
//
//	err := emailService.SendEmail(props)
//	if err != nil {
//	    log.Fatalf("Failed to send email: %v", err)
//	}
func (e *EmailServiceImpl) SendEmail(props SendEmailProps) error {
	// Establish an SMTP connection
	client, err := e.Dialer.Dial()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error establishing SMTP connection")
	}
	defer client.Close()

	// Validate the primary recipient
	if props.To == "" {
		return fmt.Errorf("receiver cannot be empty")
	}

	// Create a new email message
	message := gomail.NewMessage()
	message.SetHeader("From", e.Dialer.Username)
	message.SetHeader("To", props.To)
	message.SetHeader("Subject", props.Subject)
	message.SetBody("text/html", props.Body)
	message.SetHeader("Cc", props.Cc...)
	message.SetHeader("Bcc", props.Bcc...)

	// Attach files
	for _, attachment := range props.Attachments {
		fileData, err := base64.StdEncoding.DecodeString(attachment.Base64)
		if err != nil {
			fmt.Printf("Error decoding Base64 data for file %s: %v\n", attachment.Name, err)
			continue
		}

		message.Attach(attachment.Name, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(fileData)
			return err
		}))
	}

	// Send the email
	err = gomail.Send(client, message)

	return err
}
