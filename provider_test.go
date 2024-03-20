package emailutils

import "testing"

var testEmail = &Email{
	From:    NewPerson("Company", "noreply@domain.com"),
	To:      []*Person{NewPerson("Polo", "sespolo@gmail.com")},
	Subject: "Company contacting you",
	Body:    "Hello Polo, this Company reaching you.",
}

func TestInit(t *testing.T) {

}
