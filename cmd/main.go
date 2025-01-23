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
func getTasksFromDB() ([]string, error) {
	db := database()
	defer db.Close()
	rows, err := db.Query("SELECT my_task FROM task;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []string

	for rows.Next() {
		var task string
		err := rows.Scan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
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

func renderTemplate(w http.ResponseWriter, r *http.Request, file_name string, tasks []string) {

	t, err := template.ParseFiles(file_name)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	err = t.Execute(w, tasks)
	if err != nil {
		log.Println("Error rendering temlate")
	}
}

func mainHendler(w http.ResponseWriter, r *http.Request) {
	tasks, err := getTasksFromDB()
	if err != nil {
		log.Println("Error geting task from DB: ", err)
		return
	}
	renderTemplate(w, r, "html/main.html", tasks)
}

func taskHendler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "html/add_task.html", nil)
}

func hendlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		task, err := getTask(r)
		if err != nil {
			log.Println("Error getting data: ", err)
		}
		addTaskToDB(task)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func main() {

	http.HandleFunc("/", mainHendler)
	http.HandleFunc("/add_task", taskHendler)
	http.HandleFunc("/form", hendlerPost)

	log.Println("Server started")
	err := http.ListenAndServe(":8180", nil)
	if err != nil {
		log.Fatal(err)
	}
}
