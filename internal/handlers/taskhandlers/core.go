package taskhandlers

import "net/http"

type serviceer interface {
	Create(name string) error
	Delete(name string) error
	List() ([]string, error)
}

type handler struct {
	srv serviceer
}

func New(srv serviceer) *handler {
	return &handler{srv: srv}
}

func (h *handler) Register() {
	http.HandleFunc("/home", h.list)
	http.HandleFunc("/add-task", h.createForm)
	http.HandleFunc("/create-task", h.create)
	http.HandleFunc("/delete-task", h.delete)
}
