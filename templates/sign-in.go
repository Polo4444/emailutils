package templates

var SignIn = map[Lang]Content{
	"en": {
		Subject: "{{.Code}} is your Company sign in code",
		Body:    `<p>You have requested login Copy the code below and paste it to otp field.</p><p><b>{{.Code}}</b></p>`,
	},
	"fr": {
		Subject: "{{.Code}} est votre code de connexion",
		Body:    "<p>Nous espérons que ce message vous trouve bien.Veuillez vous connecter avec ce code:</p><p><b>{{.Code}}</b></p>",
	},
	"es": {
		Subject: "{{.Code}} es tu código de inicio de sesión",
		Body:    "<p>Esperamos que este mensaje te encuentre bien. Por favor, inicie sesión con este código:</p><p><b>{{.Code}}</b></p>",
	},
	"pt": {
		Subject: "{{.Code}} é o seu código de login",
		Body:    "<p>Esperamos que esta mensagem o encontre bem. Por favor, faça login com este código:</p><p><b>{{.Code}}</b></p>",
	},
	"de": {
		Subject: "{{.Code}} ist Ihr Anmeldecode",
		Body:    "<p>Wir hoffen, dass diese Nachricht Sie gut erreicht. Bitte melden Sie sich mit diesem Code an:</p><p><b>{{.Code}}</b></p>",
	},
	"it": {
		Subject: "{{.Code}} è il tuo codice di accesso",
		Body:    "<p>Speriamo che questo messaggio ti trovi bene. Accedi con questo codice:</p><p><b>{{.Code}}</b></p>",
	},
}
