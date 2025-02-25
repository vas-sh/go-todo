package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

//go:generate mockgen -source=core.go -destination=mocks/mocks.go -package mocks

type repoer interface {
	Create(ctx context.Context, title, description string) (models.Task, error)
	Remove(ctx context.Context, id int64) error
	List(ctx context.Context) ([]models.Task, error)
}

type srv struct {
	repo repoer
}

func New(repo repoer) *srv {
	return &srv{repo: repo}
}
