package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) ReportCompletions(ctx context.Context, userID int64) (models.CountCompletion, error) {
	return s.repo.ReportCompletions(ctx, userID)
}
