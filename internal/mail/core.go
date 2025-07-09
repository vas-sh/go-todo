package mail

import "html/template"

type srv struct {
	login    string
	password string
	host     string
	port     string
	templ    *template.Template
}

func New(login, password, host, port string) (*srv, error) {
	templ, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return nil, err
	}
	return &srv{
		login:    login,
		password: password,
		host:     host,
		port:     port,
		templ:    templ,
	}, nil
}
