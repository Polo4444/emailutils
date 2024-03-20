package emailutils

import (
	"testing"
)

var testSMTPConf = BuildSMTPConfig(
	"Host",
	465,
	"User",
	"Pass",
	false,
)

func TestSendSMTP(t *testing.T) {

	// Init
	TestInit(t)

	// We create a new SMTP provider
	p, err := GetProvider(SMTPProvider, testSMTPConf)
	if err != nil {
		t.Fatalf("Error creating SMTP provider: %s", err.Error())
	}

	// We send the email
	err = p.Send(testEmail)
	if err != nil {
		t.Fatalf("Error sending email: %s", err.Error())
	}

	t.Log("Email sent successfully\n")
}
