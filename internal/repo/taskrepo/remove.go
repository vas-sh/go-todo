package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Remove(ctx context.Context, id, userID int64) error {
	return r.db.WithContext(ctx).Model(models.Task{}).Delete(models.Task{}, "id = ? AND user_id = ?", id, userID).Error
}
