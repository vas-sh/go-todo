package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) List(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.WithContext(ctx).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, err
}
