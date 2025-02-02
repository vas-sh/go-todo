package taskhandlers

import "net/http"

func (*handler) name(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}
	task := r.FormValue("task")
	return task, nil
}
