package taskhandlers

import (
	"context"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

type serviceer interface {
	Create(ctx context.Context, title, description string) (models.Task, error)
	Remove(ctx context.Context, id int64) error
	List(ctx context.Context) ([]models.Task, error)
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

func (h *handler) Register() error {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		c.Next()
	})
	r.GET(h.homePath, h.home)
	r.GET("/add-task", h.addTask)
	r.POST("/create-task", h.create)
	r.POST("/delete-task", h.remove)
	tasksRouter := r.Group("/api/tasks")
	tasksRouter.GET("", h.homeAPI)
	tasksRouter.POST("", h.createAPI)
	tasksRouter.DELETE("", h.removeAPI)
	tasksRouter.OPTIONS("", func(_ *gin.Context) {})
	return r.Run()
}
