package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, file_name string, tasks []string) {
	t, err := template.ParseFiles(file_name)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	err = t.Execute(w, tasks)
	if err != nil {
		log.Println("Error with writing task: ", err)
	}
}
