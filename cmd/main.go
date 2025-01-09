package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/lib/pq"
)

func database() *sql.DB {
	DSN := "host=localhost user=user password=1111 dbname=test port=5432 sslmode=disable TimeZone=Europe/Kiev"

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func hendler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("html", "bace.html")

	t, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, "Error parse file ", http.StatusInternalServerError)

		log.Println("Error parse file ", err)
	}

	err = t.Execute(w, "Hello")
	if err != nil {
		http.Error(w, "Error rendering templete ", http.StatusInternalServerError)
		log.Println("Error rendering templete: ", err)
	}
}

func main() {
	http.HandleFunc("/", hendler)
	db := database()
	defer db.Close()

	log.Println("Server started")
	err := http.ListenAndServe(":8180", nil)
	if err != nil {
		log.Fatal(err)
	}
}
