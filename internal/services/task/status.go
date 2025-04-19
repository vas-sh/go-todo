package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Statuses(ctx context.Context, userID, taskID int64) ([]models.TaskStatus, error) {
	return s.repo.Statuses(ctx, userID, taskID)
}
