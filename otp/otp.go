package otp

import (
	"bytes"
	"text/template"
	"time"

	"github.com/Polo4444/emailutils"
	"github.com/google/uuid"
)

type Store interface {
	Register(code, purpose, entity string) (*EmailCode, error)
	List() ([]EmailCode, error)
	Get(code, purpose string) (*EmailCode, error)
	Remove(code, purpose string) error
}

type OTP struct {
	CM         Store
	ExpireIn   time.Duration
	CodeLength int
}

// NewOTP creates a new OTP
func NewOTP(cm Store, expireIn time.Duration, codeLength int) *OTP {
	return &OTP{
		CM:         cm,
		ExpireIn:   expireIn,
		CodeLength: codeLength,
	}
}

// Verify verifies the code and removes it from the list of codes if verified
func (o *OTP) Verify(code, purpose, entity string) bool {

	// Get code
	em, err := o.CM.Get(code, purpose)
	if err != nil {
		return false
	}

	if em.checkIfCodeValid(o.ExpireIn) && em.Entity == entity {
		go o.CM.Remove(code, purpose) // We remove the code
		return true
	}

	return false
}

// Send sends an email with the code and registers it.
// It will assume your email subject and body has a template with the following keys:
// - {{.Code}}: the code to send
func (o *OTP) Send(
	email *emailutils.Email,
	provider emailutils.Provider,
	purpose, entity string,
) error {

	// Generate code
	code := emailutils.RandomNum(o.CodeLength)

	// Build Subject
	subject, err := NewBuilder(&email.Subject, map[string]string{"Code": code}).Build()
	if err != nil {
		return err
	}
	email.Subject = subject

	// Build Body
	body, err := NewBuilder(&email.Body, map[string]string{"Code": code}).Build()
	if err != nil {
		return err
	}
	email.Body = body

	// Send email
	err = provider.Send(email)
	if err != nil {
		return err
	}

	// Register code
	_, err = o.CM.Register(code, purpose, entity)
	if err != nil {
		return err
	}

	return nil
}

// ProcessEmailsCode processes the codes by spawning a ticker of teh specified interval.
// It checks the codes and removes the expired ones from the list of codes
// It returns a function to stop the process.
func (o *OTP) ProcessEmailsCode(interval time.Duration) func() {

	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for range ticker.C {

			// grab codes
			Codes, err := o.CM.List()
			if err != nil {
				continue
			}

			// We check list of Codes
			for _, v := range Codes {
				if !v.checkIfCodeValid(o.ExpireIn) {
					o.CM.Remove(v.Code, v.Purpose)
				}
			}
		}
	}()

	return ticker.Stop
}

//--------------------------------------------------------------------------------------//
//                                       Builder                                        //
//--------------------------------------------------------------------------------------//

type Builder struct {
	Data     *string
	MetaData map[string]string
}

func NewBuilder(data *string, metaData map[string]string) *Builder {
	return &Builder{Data: data, MetaData: metaData}
}

func (d *Builder) Build() (string, error) {

	// We parse subject
	tmpl := template.New(uuid.Must(uuid.NewRandom()).String())
	tmpl = tmpl.Option("missingkey=zero")
	tmpl, err := tmpl.Parse(*d.Data)
	if err != nil {
		return "", err
	}

	data := bytes.NewBuffer(nil)
	err = tmpl.Execute(data, d.MetaData)
	if err != nil {
		return "", err
	}

	return data.String(), nil
}
