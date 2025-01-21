package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func database() *sql.DB {
	DSN := "host=localhost user=vas password=2222 dbname=test_db port=5432 sslmode=disable TimeZone=Europe/Kiev"

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getTask(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}
	task := r.FormValue("task")
	return task, err

}

func addTaskToDB(task string) {
	db := database()
	defer db.Close()
	f := `INSERT INTO task (my_task) VALUES ($1)`
	if task == "" {
		log.Println("Value can't be empty")
		return
	}
	_, err := db.Exec(f, task)
	if err != nil {
		log.Printf("Error inserting value to db: %s", err)
		return
	}
	log.Println("Task secsesfuly added")
}

func hendlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
	}

	task, err := getTask(r)
	if err != nil {
		log.Printf("Error getting data from form: %s", err)
	}
	addTaskToDB(task)
}

func hendlerGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("html/bace.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Println("Error rendering template:", err)
		}
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			hendlerGet(w, r)
		} else if r.Method == http.MethodPost {
			hendlerPost(w, r)
		} else {
			log.Println("Method is not defineted")
		}
	})

	log.Println("Server started")
	err := http.ListenAndServe(":8180", nil)
	if err != nil {
		log.Fatal(err)
	}
}
