package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/vas-sh/todo/internal/config"
	"github.com/vas-sh/todo/internal/db"
	"github.com/vas-sh/todo/internal/handlers"
	"github.com/vas-sh/todo/internal/handlers/taskhandlers"
	"github.com/vas-sh/todo/internal/handlers/userhandlers"
	"github.com/vas-sh/todo/internal/repo/taskrepo"
	"github.com/vas-sh/todo/internal/repo/userrepo"
	"github.com/vas-sh/todo/internal/services/jwttoken"
	"github.com/vas-sh/todo/internal/services/task"
	"github.com/vas-sh/todo/internal/services/user"
)

func main() {
	dns := "host=localhost user=todouser password=2222 dbname=tododb port=5432 sslmode=disable TimeZone=Europe/Kiev"
	databace, err := db.New(dns)
	if err != nil {
		log.Fatal(err)
	}
	databace = databace.Debug()
	err = db.Migrate(databace)
	if err != nil {
		log.Fatal(err)
	}
	taskRepo := taskrepo.New(databace)
	taskSrv := task.New(taskRepo)

	userRepo := userrepo.New(databace)
	userSrv := user.New(userRepo)

	userFetcher := jwttoken.New(config.Config.SecretJWT)
	server := handlers.New(userFetcher)
	anonRouter := server.AnonRouter()
	authRouter := server.AuthRouter()
	taskhandlers.New(taskSrv).Register(authRouter)
	userhandlers.New(userSrv, config.Config.SecretJWT).Register(anonRouter, authRouter)

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
