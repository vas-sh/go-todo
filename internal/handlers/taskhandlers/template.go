package taskhandlers

import (
	"html/template"
	"net/http"
)

func (*handler) renderTemplate(w http.ResponseWriter, file_name string, tasks []string) {
	t, err := template.ParseFiles(file_name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
