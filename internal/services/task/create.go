package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Create(ctx context.Context, title, description string) (models.Task, error) {
	if title == "" {
		return models.Task{}, models.ErrValueEmpty
	}
	return s.repo.Create(ctx, title, description)
}
