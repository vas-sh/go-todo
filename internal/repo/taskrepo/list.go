package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) List(ctx context.Context) ([]string, error) {
	var tasks []models.Task
	err := r.db.WithContext(ctx).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	var names []string
	for i := range tasks {
		names = append(names, tasks[i].Title)
	}
	return names, err
}
