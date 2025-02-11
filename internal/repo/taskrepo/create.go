package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Create(ctx context.Context, title, description string) error {
	return r.db.WithContext(ctx).Create(&models.Task{
		Title:       title,
		Status:      models.NewStatus,
		Description: description,
	}).Error
}
