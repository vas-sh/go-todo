package taskhandlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

type serviceer interface {
	Create(ctx context.Context, body models.Task) (models.Task, error)
	Remove(ctx context.Context, id, userID int64) error
	List(ctx context.Context, userID int64) ([]models.Task, error)
	Update(ctx context.Context, body models.Task, userID, taskID int64) error
	Statuses(ctx context.Context, userID, taskID int64) ([]models.TaskStatus, error)
	ReportStatuses(ctx context.Context, userID int64) ([]models.CountStatus, error)
	ReportCompletions(ctx context.Context, userID int64) (models.CountCompletion, error)
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
	tasksRouter.OPTIONS("/:id", func(_ *gin.Context) {})
	tasksRouter.OPTIONS("/:id/statuses", func(_ *gin.Context) {})
	tasksRouter.PUT("/:id", h.update)
	tasksRouter.GET("/:id/statuses", h.statuses)
	tasksRouter.GET("/report-statuses", h.reportStatuses)
	tasksRouter.GET("/report-completions", h.reportCompletions)
}
