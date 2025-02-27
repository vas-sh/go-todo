package taskhandlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

type serviceer interface {
	Create(ctx context.Context, title, description string) (models.Task, error)
	Remove(ctx context.Context, id int64) error
	List(ctx context.Context) ([]models.Task, error)
}

type handler struct {
	srv serviceer
}

func New(srv serviceer) *handler {
	return &handler{
		srv: srv,
	}
}

func (h *handler) Register() error {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Next()
	})
	tasksRouter := r.Group("/api/tasks")
	tasksRouter.GET("", h.list)
	tasksRouter.POST("", h.create)
	tasksRouter.DELETE("", h.remove)
	tasksRouter.OPTIONS("", func(_ *gin.Context) {})
	return r.Run()
}
