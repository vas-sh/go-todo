package user

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

type repoer interface {
	Create(ctx context.Context, user *models.User) error
}

type srv struct {
	repo repoer
}

func New(repo repoer) *srv {
	return &srv{repo: repo}
}
