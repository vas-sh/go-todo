package mail

import (
	"net/smtp"
)

func (s *srv) Send(to, subject, body string) error {
	msg := "From: " + s.login + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	return smtp.SendMail(s.host+":"+s.port,
		smtp.PlainAuth("", s.login, s.password, s.host),
		s.login, []string{to}, []byte(msg))
}
