package  emailutils

import (
	"testing"
)

var testMailgunConf = BuildMailgunConfig(
	"Domain",
	"APIKey",
)

func TestSendMailgun(t *testing.T) {

	// Init
	TestInit(t)

	// We create a new SMTP provider
	p, err := GetProvider(MailgunProvider, testMailgunConf)
	if err != nil {
		t.Fatalf("Error creating Mailgun provider: %s", err.Error())
	}

	// We send the email
	err = p.Send(testEmail)
	if err != nil {
		t.Fatalf("Error sending email: %s", err.Error())
	}

	t.Log("Email sent successfully\n")
}
