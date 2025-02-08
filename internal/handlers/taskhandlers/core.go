package taskhandlers

import (
	"context"
	"html/template"
	"net/http"
)

type serviceer interface {
	Create(ctx context.Context, title string) error
	Remove(ctx context.Context, title string) error
	List(ctx context.Context) ([]string, error)
}

type handler struct {
	srv                serviceer
	homePath           string
	createFormTemplate *template.Template
	homeTemplete       *template.Template
}

func New(srv serviceer) (*handler, error) {
	createFormTemplate, err := template.ParseFiles("html/add-task.html")
	if err != nil {
		return nil, err
	}
	homeTemplate, err := template.ParseFiles("html/home.html")
	if err != nil {
		return nil, err
	}
	return &handler{
		srv:                srv,
		homePath:           "/home",
		createFormTemplate: createFormTemplate,
		homeTemplete:       homeTemplate,
	}, nil
}

func (h *handler) Register() {
	http.HandleFunc(h.homePath, h.list)
	http.HandleFunc("/add-task", h.createForm)
	http.HandleFunc("/create-task", h.create)
	http.HandleFunc("/delete-task", h.remove)
}
