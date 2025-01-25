package handlers

import (
	"log"
	"net/http"

	"github.com/vas-sh/todo/databace"
)

func GetTaskFromForm(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}
	task := r.FormValue("task")
	return task, err

}

func HandlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		task, err := GetTaskFromForm(r)
		if err != nil {
			log.Println("Error getting data: ", err)
		}
		databace.AddTaskToDB(task)
		http.Redirect(w, r, "/main", http.StatusSeeOther)
	}

}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := databace.GetTasksFromDB()
	if err != nil {
		log.Println("Error geting task from DB: ", err)
		return
	}
	RenderTemplate(w, r, "html/main.html", tasks)
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, "html/add_task.html", nil)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form: ", err)
		}
	}
	task := r.FormValue("task")
	databace.DeleteTask(task)
	http.Redirect(w, r, "/main", http.StatusSeeOther)
}
