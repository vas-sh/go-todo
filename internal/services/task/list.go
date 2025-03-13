package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) List(ctx context.Context, userID int64) ([]models.Task, error) {
	return s.repo.List(ctx, userID)
}
