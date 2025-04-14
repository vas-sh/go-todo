package taskhandlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

type serviceer interface {
	Create(ctx context.Context, title, description string, userID int64) (models.Task, error)
	Remove(ctx context.Context, id, userID int64) error
	List(ctx context.Context, userID int64) ([]models.Task, error)
	Update(ctx context.Context, body models.Task, userID int64, taskID int64) error
}

type handler struct {
	srv serviceer
}

func New(srv serviceer) *handler {
	return &handler{
		srv: srv,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	tasksRouter := router.Group("tasks")
	tasksRouter.GET("", h.list)
	tasksRouter.POST("", h.create)
	tasksRouter.DELETE("/:id", h.remove)
	tasksRouter.OPTIONS("", func(_ *gin.Context) {})
	tasksRouter.PUT("/:id", h.update)
}
