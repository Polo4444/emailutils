package emailutils

import (
	"io"

	"gopkg.in/gomail.v2"
)

type smtpProvider struct {
	server  string
	port    int
	user    string
	pass    string
	skipTls bool
}

func BuildSMTPConfig(server string, port int, user, pass string, skipTls bool) ProviderConfig {
	return ProviderConfig{
		"server":  server,
		"port":    port,
		"user":    user,
		"pass":    pass,
		"skipTls": skipTls,
	}
}

func newSMTPProvider(config ProviderConfig) (*smtpProvider, error) {

	return &smtpProvider{
		server:  config["server"].(string),
		port:    config["port"].(int),
		user:    config["user"].(string),
		pass:    config["pass"].(string),
		skipTls: config["skipTls"].(bool),
	}, nil
}

// Send sends an email using gomail
func (s *smtpProvider) Send(e *Email) error {

	m := gomail.NewMessage()

	// Validations
	err := e.Validate()
	if err != nil {
		return err
	}

	// Tos
	tos := make([]string, len(e.To))
	for i, to := range e.To {
		tos[i] = m.FormatAddress(to.Email, to.Name)
	}

	// Ccs
	ccs := make([]string, len(e.Cc))
	for i, cc := range e.Cc {
		ccs[i] = m.FormatAddress(cc.Email, cc.Name)
	}

	// Bccs
	bccs := make([]string, len(e.Bcc))
	for i, bcc := range e.Bcc {
		bccs[i] = m.FormatAddress(bcc.Email, bcc.Name)
	}

	m.SetHeaders(map[string][]string{
		"Subject": {e.Subject},
		"From":    {m.FormatAddress(e.From.Email, e.From.Name)},
		"To":      tos,
		"Cc":      ccs,
		"Bcc":     bccs,
	})
	m.SetBody("text/html", e.Body)

	// Attachments
	for _, a := range e.Attachments {
		m.Attach(a.Name, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(a.Data)
			return err
		}))
	}

	d := gomail.NewDialer(s.server, s.port, s.user, s.pass)
	// d.SSL = true
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: s.skipTls}

	return d.DialAndSend(m)
}

func (s *smtpProvider) SupportAttachments() bool {
	return true
}
