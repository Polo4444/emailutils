package emailutils

import (
	"testing"
)

var testSendgridConfig = BuildSendgridConfig(
	"APIKey",
)

func TestSendSendgrid(t *testing.T) {

	// Init
	TestInit(t)

	// We create a new SMTP provider
	p, err := GetProvider(SendgridProvider, testSendgridConfig)
	if err != nil {
		t.Fatalf("Error creating Sendgrid provider: %s", err.Error())
	}

	// We send the email
	err = p.Send(testEmail)
	if err != nil {
		t.Fatalf("Error sending email: %s", err.Error())
	}

	t.Log("Email sent successfully\n")
}
