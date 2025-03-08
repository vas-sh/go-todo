package user

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

type repoer interface {
	Create(ctx context.Context, user *models.User) error
	Remove(ctx context.Context, id int64) error
	GetByEmail(ctx context.Context, email string) (user models.User, err error)
}

type srv struct {
	repo repoer
}

func New(repo repoer) *srv {
	return &srv{repo: repo}
}
