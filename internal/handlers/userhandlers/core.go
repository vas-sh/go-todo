package userhandlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

type serviceer interface {
	SignUp(ctx context.Context, body models.CreateUserBody) (*models.User, error)
	Remove(ctx context.Context, id int64) error
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
	usersRouter := router.Group("users")
	usersRouter.POST("/sign-up", h.signUp)
	usersRouter.DELETE("", h.remove)
	usersRouter.OPTIONS("", func(_ *gin.Context) {})
	usersRouter.OPTIONS("/:r", func(_ *gin.Context) {})
}
