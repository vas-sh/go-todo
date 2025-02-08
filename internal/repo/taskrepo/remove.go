package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Remove(ctx context.Context, title string) error {
	return r.db.WithContext(ctx).Model(models.Task{}).Delete(models.Task{}, "title = ?", title).Error
}
