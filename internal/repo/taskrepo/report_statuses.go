package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) ReportStatuses(ctx context.Context, userID int64) ([]models.CountStatus, error) {
	var statuses []models.CountStatus
	err := r.db.WithContext(ctx).
		Model(models.Task{}).
		Select("status, COUNT(1) as count").
		Where("user_id = ?", userID).
		Group("status").
		Find(&statuses).Error
	if err != nil {
		return nil, err
	}
	return statuses, nil
}
