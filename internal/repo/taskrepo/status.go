package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Statuses(ctx context.Context, userID, taskID int64) ([]models.TaskStatus, error) {
	var statuses []models.TaskStatus
	err := r.db.WithContext(ctx).
		Where("task_id = (SELECT t.id FROM tasks AS t WHERE t.user_id = ? AND t.id = ?)", userID, taskID).
		Order("id").Find(&statuses).Error
	return statuses, err
}
