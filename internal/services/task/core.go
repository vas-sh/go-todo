package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

//go:generate mockgen -source=core.go -destination=mocks/mocks.go -package mocks

type repoer interface {
	Create(ctx context.Context, res *models.Task) error
	Remove(ctx context.Context, id, userID int64) error
	List(ctx context.Context, userID int64) ([]models.Task, error)
	Update(ctx context.Context, body models.Task, userID, taskID int64) error
	Statuses(ctx context.Context, userID, taskID int64) ([]models.TaskStatus, error)
	ReportStatuses(ctx context.Context, userID int64) (models.CountStatus, error)
	ReportCompletions(ctx context.Context, userID int64) (models.CountCompletion, error)
}

type srv struct {
	repo repoer
}

func New(repo repoer) *srv {
	return &srv{repo: repo}
}
