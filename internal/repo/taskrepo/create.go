package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Create(ctx context.Context, res *models.Task) error {
	err := r.db.WithContext(ctx).Create(res).Error
	return err
}
