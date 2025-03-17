package mail

type srv struct {
	login    string
	password string
	host     string
	port     string
}

func New(login, password, host, port string) *srv {
	return &srv{
		login:    login,
		password: password,
		host:     host,
		port:     port,
	}
}
