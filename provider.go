package emailutils

import (
	"fmt"
)

type ProviderEntity string
type ProviderConfig map[string]interface{}

const (
	SMTPProvider ProviderEntity = "SMTP"
	// TODO: implement breo provider. // BrevoProvider    ProviderEntity = "BREVO"
	MailgunProvider  ProviderEntity = "MAILGUN"
	SendgridProvider ProviderEntity = "SENDGRID"
)

// Errors
var (
	ErrProviderNotSupported = fmt.Errorf("DNS Provider not supported")
	ErrNoFrom               = fmt.Errorf("from not provided")
	ErrNoRecipients         = fmt.Errorf("no recipients provided")
	ErrNoSubject            = fmt.Errorf("no subject provided")
	ErrNoBody               = fmt.Errorf("no body provided")
)

type Person struct {
	Name  string
	Email string
}

func NewPerson(name, email string) *Person {
	return &Person{Name: name, Email: email}
}

func (p *Person) String() string {
	return fmt.Sprintf("%s <%s>", p.Name, p.Email)
}

func (p *Person) Validate() error {

	if p.Email == "" {
		return fmt.Errorf("no email provided")
	}

	return nil
}

type Attachment struct {
	Name        string
	ContentType string
	Data        []byte
}

func NewAttachment(name, contentType string, data []byte) *Attachment {
	return &Attachment{Name: name, ContentType: contentType, Data: data}
}

func (a *Attachment) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("no name provided")
	}

	if len(a.Data) == 0 {
		return fmt.Errorf("no data provided")
	}

	return nil
}

// Define Email struct
type Email struct {
	Subject     string
	From        *Person
	To          []*Person
	Cc          []*Person
	Bcc         []*Person
	Body        string
	Attachments []*Attachment
}

func (e *Email) Validate() error {

	if e.From.Validate() != nil {
		return ErrNoFrom
	}

	for _, to := range e.To {
		if to.Validate() != nil {
			return ErrNoRecipients
		}
	}

	if e.Subject == "" {
		return ErrNoSubject
	}

	if e.Body == "" {
		return ErrNoBody
	}

	return nil
}

// Define the Provider interface
type Provider interface {
	Send(e *Email) error
	SupportAttachments() bool
}

// GetProvider returns an email provider
func GetProvider(entity ProviderEntity, config ProviderConfig) (Provider, error) {

	var err error = nil
	var emailProvider Provider = nil

	switch entity {
	case SMTPProvider:
		emailProvider, err = newSMTPProvider(config)
	case MailgunProvider:
		emailProvider, err = newMailgunProvider(config)
	case SendgridProvider:
		emailProvider, err = newSendgridProvider(config)
	default:
		err = ErrProviderNotSupported
	}

	return emailProvider, err
}
