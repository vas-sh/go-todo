package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Update(ctx context.Context, body models.Task, userID, taskID int64) error {
	return s.repo.Update(ctx, body, userID, taskID)
}
