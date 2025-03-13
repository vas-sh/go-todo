package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Create(ctx context.Context, title, description string, userID int64) (models.Task, error) {
	if title == "" {
		return models.Task{}, models.ErrValueEmpty
	}
	return s.repo.Create(ctx, title, description, userID)
}
