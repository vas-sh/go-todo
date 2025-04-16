package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Create(ctx context.Context, body models.Task) (models.Task, error) {
	if body.Title == "" {
		return models.Task{}, models.ErrValueEmpty
	}
	body.ID = 0
	err := s.repo.Create(ctx, &body)
	return body, err
}
