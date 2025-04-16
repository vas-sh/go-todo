package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Update(ctx context.Context, body models.Task, userID, taskID int64) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND id = ?", userID, taskID).
		Model(models.Task{}).
		Updates(body).Error
}
