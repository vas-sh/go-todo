package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/vas-sh/todo/internal/db"
	"github.com/vas-sh/todo/internal/handlers"
	"github.com/vas-sh/todo/internal/handlers/taskhandlers"
	"github.com/vas-sh/todo/internal/handlers/userhandlers"
	"github.com/vas-sh/todo/internal/repo/taskrepo"
	"github.com/vas-sh/todo/internal/repo/userrepo"
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

	server := handlers.New()
	router := server.Router()
	taskhandlers.New(taskSrv).Register(router)
	userhandlers.New(userSrv).Register(router)

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
