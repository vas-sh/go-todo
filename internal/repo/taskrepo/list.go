package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) List(ctx context.Context, userID int64) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order(`
		CASE status WHEN 'inProgress' THEN 0
					WHEN 'new' THEN 1
					WHEN 'done' THEN 2
					ELSE 3
		END,
		id DESC
	`).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, err
}
