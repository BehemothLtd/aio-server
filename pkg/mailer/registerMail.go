package mailer

import (
	"bytes"
	"text/template"
)

type RegisterEmailData struct {
	Name    string
	Message string
}

const RegisterTmpl = `
<html>
	<body>
			<p>Dear {{.Name}},</p>
			<p>{{.Message}}</p>
	</body>
</html>
`

func (emailData *RegisterEmailData) Send(to string, subject string) error {
	emailMessage := NewEmailMessage(to, subject)

	tmpl := template.Must(template.New("RegisterTemplate").Parse(RegisterTmpl))

	var body bytes.Buffer

	if err := tmpl.Execute(&body, emailData); err != nil {
		return err
	}

	emailMessage.Body = body.String()

	if err := Send(*emailMessage); err != nil {
		return err
	}

	return nil
}
