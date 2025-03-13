package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) List(ctx context.Context, userID int64) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, err
}
