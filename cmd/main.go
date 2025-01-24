package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vas-sh/todo/handlers"
)

func main() {

	http.HandleFunc("/main", handlers.MainHandler)
	http.HandleFunc("/add_task", handlers.TaskHandler)
	http.HandleFunc("/form", handlers.HandlerPost)

	log.Println("Server started")
	err := http.ListenAndServe(":8180", nil)
	if err != nil {
		log.Fatal(err)
	}
}
