package templates

var PasswordForgotten = map[Lang]Content{
	"en": {
		Subject: "{{.Code}} is your password reset code",
		Body:    `<p>You have requested a password reset. Copy the code below and paste it to the password reset field.</p><p><b>{{.Code}}</b></p>`,
	},
	"fr": {
		Subject: "{{.Code}} est votre code de réinitialisation de mot de passe",
		Body:    "<p>Vous avez demandé une réinitialisation de mot de passe. Copiez le code ci-dessous et collez-le dans le champ de réinitialisation de mot de passe.</p><p><b>{{.Code}}</b></p>",
	},
	"es": {
		Subject: "{{.Code}} es tu código de restablecimiento de contraseña",
		Body:    "<p>Has solicitado un restablecimiento de contraseña. Copia el código a continuación y pégalo en el campo de restablecimiento de contraseña.</p><p><b>{{.Code}}</b></p>",
	},
	"pt": {
		Subject: "{{.Code}} é o seu código de redefinição de senha",
		Body:    "<p>Você solicitou uma redefinição de senha. Copie o código abaixo e cole-o no campo de redefinição de senha.</p><p><b>{{.Code}}</b></p>",
	},
	"de": {
		Subject: "{{.Code}} ist Ihr Passwort-Reset-Code",
		Body:    "<p>Sie haben einen Passwort-Reset angefordert. Kopieren Sie den Code unten und fügen Sie ihn in das Passwort-Reset-Feld ein.</p><p><b>{{.Code}}</b></p>",
	},
	"it": {
		Subject: "{{.Code}} è il tuo codice di reimpostazione della password",
		Body:    "<p>Hai richiesto un ripristino della password. Copia il codice di seguito e incollalo nel campo di ripristino della password.</p><p><b>{{.Code}}</b></p>",
	},
}
