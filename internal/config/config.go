package config

import "os"

var Config = struct {
	SecretJWT string
}{
	SecretJWT: os.Getenv("SECRET_JWT"),
}
