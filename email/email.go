package email

import (
	"encoding/base64"
	"fmt"
	"io"

	"gopkg.in/gomail.v2"
)

type EmailAttachment struct {
	Base64 string
	Name   string
}

type SendEmailProps struct {
	To          string
	Subject     string
	Body        string
	Attachments []EmailAttachment
	Cc          []string
	Bcc         []string
}

type EmailService interface {
	SendEmail(props SendEmailProps) error
}

type EmailServiceImpl struct {
	Dialer *gomail.Dialer
}

func MakeEmailServiceImpl(
	user string, password string, host string, port int,
) *EmailServiceImpl {
	dialer := gomail.NewDialer(host, port, user, password)

	return &EmailServiceImpl{
		Dialer: dialer,
	}
}

func (e *EmailServiceImpl) SendEmail(props SendEmailProps) error {
	client, err := e.Dialer.Dial()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error establishin smpt connection")
	}
	defer client.Close()

	if props.To == "" {
		return fmt.Errorf("receiver cannot be empty")
	}

	message := gomail.NewMessage()
	message.SetHeader("From", e.Dialer.Username)
	message.SetHeader("To", props.To)
	message.SetHeader("Subject", props.Subject)
	message.SetHeader("Cc", props.Cc...)
	message.SetHeader("Bcc", props.Bcc...)
	message.SetBody("text/html", props.Body)

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

	err = gomail.Send(client, message)

	return err
}
