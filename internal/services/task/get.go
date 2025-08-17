package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) GetTask(ctx context.Context, userID, offset int64) (models.Task, bool, error) {
	return s.repo.GetTask(ctx, userID, offset)
}
