package taskhandlers

import "net/http"

type taskBody struct {
	title       string
	description string
}

func (*handler) body(r *http.Request) (taskBody, error) {
	err := r.ParseForm()
	if err != nil {
		return taskBody{}, err
	}
	return taskBody{
		title:       r.FormValue("title"),
		description: r.FormValue("description"),
	}, nil
}
