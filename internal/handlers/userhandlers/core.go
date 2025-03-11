package userhandlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

type serviceer interface {
	SignUp(ctx context.Context, body models.CreateUserBody) (*models.User, error)
	Remove(ctx context.Context, id int64) error
	Login(ctx context.Context, email, password string) (models.User, error)
}

type handler struct {
	srv       serviceer
	secretJWT string
}

func New(srv serviceer, secretJWT string) *handler {
	return &handler{
		srv:       srv,
		secretJWT: secretJWT,
	}
}

func (h *handler) Register(anonRouter *gin.RouterGroup, authRouter *gin.RouterGroup) {
	usersAnonRouter := anonRouter.Group("users")
	usersAnonRouter.POST("/sign-up", h.signUp)
	usersAnonRouter.OPTIONS("", func(_ *gin.Context) {})
	usersAnonRouter.OPTIONS("/:r", func(_ *gin.Context) {})
	usersAnonRouter.POST("/login", h.login)

	usersAuthRouter := authRouter.Group("users")
	usersAuthRouter.DELETE("", h.remove)
}
