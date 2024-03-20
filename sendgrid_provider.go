package emailutils

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendgridProvider struct {
	apiKey string
}

func BuildSendgridConfig(apiKey string) ProviderConfig {
	return ProviderConfig{
		"apiKey": apiKey,
	}
}

func newSendgridProvider(config ProviderConfig) (*sendgridProvider, error) {

	return &sendgridProvider{
		apiKey: config["apiKey"].(string),
	}, nil
}

func (s *sendgridProvider) formatAddress(person *Person) *mail.Email {
	return mail.NewEmail(person.Name, person.Email)
}

func (s *sendgridProvider) Send(e *Email) error {

	// Validations
	err := e.Validate()
	if err != nil {
		return err
	}

	// Create an instance of the Client
	client := sendgrid.NewSendClient(s.apiKey)

	// We build the message
	message := mail.NewSingleEmail(
		s.formatAddress(e.From),
		e.Subject,
		s.formatAddress(e.To[0]),
		"",
		e.Body,
	)

	// Add recipients
	personalizations := mail.NewPersonalization()
	for _, recipient := range e.To[1:] {
		personalizations.AddTos(s.formatAddress(recipient))
	}

	// Handle CCs
	for _, cc := range e.Cc {
		personalizations.AddCCs(s.formatAddress(cc))
	}

	// Handle BCCs
	for _, bcc := range e.Bcc {
		personalizations.AddBCCs(s.formatAddress(bcc))
	}

	if len(personalizations.To) > 0 || len(personalizations.CC) > 0 || len(personalizations.BCC) > 0 {
		message.AddPersonalizations(personalizations)
	}

	// Handle attachments
	for _, attachment := range e.Attachments {
		message.AddAttachment(&mail.Attachment{
			Content:  base64.StdEncoding.EncodeToString(attachment.Data),
			Name:     attachment.Name,
			Filename: attachment.Name,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, err := client.SendWithContext(ctx, message)

	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf(resp.Body)
	}
	return err
}

func (m *sendgridProvider) SupportAttachments() bool {
	return true
}
