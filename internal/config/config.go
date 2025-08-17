package config

import "os"

var Config = struct {
	SecretJWT    string
	MailLogin    string
	MailPassword string
	MailHost     string
	MailPort     string
	DB           string
	TgBotToken   string
}{
	SecretJWT:    os.Getenv("SECRET_JWT"),
	MailLogin:    os.Getenv("MAIL_LOGIN"),
	MailPassword: os.Getenv("MAIL_PASSWORD"),
	MailHost:     os.Getenv("MAIL_HOST"),
	MailPort:     os.Getenv("MAIL_PORT"),
	DB:           os.Getenv("DB_URL"),
	TgBotToken:   os.Getenv("TG_BOT_TOKEN"),
}
