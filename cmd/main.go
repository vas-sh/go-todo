package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vas-sh/todo/internal/db"
	"github.com/vas-sh/todo/internal/handlers/taskhandlers"
	"github.com/vas-sh/todo/internal/repo/taskrepo"
	"github.com/vas-sh/todo/internal/services/task"
)

func main() {
	dns := "host=localhost user=vas password=2222 dbname=test_db port=5432 sslmode=disable TimeZone=Europe/Kiev"
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
	taskHandlers, err := taskhandlers.New(taskSrv)
	if err != nil {
		log.Fatal(err)
	}
	taskHandlers.Register()
	log.Println("Server started")
	err = http.ListenAndServe(":8180", nil)
	if err != nil {
		log.Fatal(err)
	}
}
