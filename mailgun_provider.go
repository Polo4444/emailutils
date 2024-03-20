package emailutils

import (
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type mailgunProvider struct {
	domain string
	apiKey string
}

func BuildMailgunConfig(domain, apiKey string) ProviderConfig {
	return ProviderConfig{
		"domain": domain,
		"apiKey": apiKey,
	}
}

func newMailgunProvider(config ProviderConfig) (*mailgunProvider, error) {

	return &mailgunProvider{
		domain: config["domain"].(string),
		apiKey: config["apiKey"].(string),
	}, nil
}

func (m *mailgunProvider) formatAddress(person *Person) string {
	return fmt.Sprintf("%s <%s>", person.Name, person.Email)
}

func (m *mailgunProvider) Send(e *Email) error {

	// Validations
	err := e.Validate()
	if err != nil {
		return err
	}

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(m.domain, m.apiKey)

	//When you have an EU-domain, you must specify the endpoint:
	//mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(m.formatAddress(e.From), e.Subject, "")
	message.SetHtml(e.Body)

	// Tos
	for _, to := range e.To {
		message.AddRecipient(m.formatAddress(to))
	}

	// Ccs
	for _, cc := range e.Cc {
		message.AddCC(m.formatAddress(cc))
	}

	// Bccs
	for _, bcc := range e.Bcc {
		message.AddBCC(m.formatAddress(bcc))
	}

	// Attachments
	for _, a := range e.Attachments {
		message.AddBufferAttachment(a.Name, a.Data)
	}

	// Send the message with a 20 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// Send the message with a 10 second timeout
	_, _, err = mg.Send(ctx, message)

	if err != nil {
		return err
	}

	return nil
}

func (m *mailgunProvider) SupportAttachments() bool {
	return true
}
