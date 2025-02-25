package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Create(ctx context.Context, title, description string) (models.Task, error) {
	res := models.Task{
		Title:       title,
		Status:      models.NewStatus,
		Description: description,
	}
	err := r.db.WithContext(ctx).Create(&res).Error
	return res, err
}
