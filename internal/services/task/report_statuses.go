package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) ReportStatuses(ctx context.Context, userID int64) ([]models.CountStatus, error) {
	return s.repo.ReportStatuses(ctx, userID)
}
