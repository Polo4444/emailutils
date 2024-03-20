package otp

import (
	"time"
)

type EmailCode struct {
	CreatedAt time.Time `json:"CreatedAt" bson:"CreatedAt"`
	Purpose   string    `json:"Purpose" bson:"Purpose"`
	Code      string    `json:"Code" bson:"Code"`
	Entity    string    `json:"Entity" bson:"Entity"` // something identifying the Entity, maybe the user email address
}

// checkIfCodeValid checks if the code has expired. It returns true if the code is still valid, false otherwise
func (c *EmailCode) checkIfCodeValid(expireIn time.Duration) bool {
	return c.CreatedAt.UTC().Add(expireIn).After(time.Now().UTC())
}

/*


// SendEmail sends an email
func SendNoReplyEmail(subject, body string, to []*Person, attachments ...*Attachment) error {

	e := &Email{
		From:        NewPerson(NoReplyName, NoReplyEmail),
		To:          to,
		Subject:     subject,
		Body:        body,
		Attachments: attachments,
	}

	// os.WriteFile("email.html", []byte(body), 0644)

	// We create a new SMTP provider
	p, err := GetProvider(SendgridProvider, BuildSendgridConfig(
		"PConfig.Sendgrid.APIKey",
	))
	if err != nil {
		return err
	}

	// We send the email
	err = p.Send(e)
	if err != nil {
		return err
	}

	return nil
}
*/
