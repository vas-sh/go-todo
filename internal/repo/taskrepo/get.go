package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) GetTask(ctx context.Context, userID, offset int64) (models.Task, bool, error) {
	q := `
		SELECT *
		FROM tasks
		WHERE user_id = ?
		ORDER BY status
		LIMIT 2 OFFSET ?
	`
	var tasks []models.Task
	err := r.db.WithContext(ctx).Raw(q, userID, offset).Scan(&tasks).Error
	if err != nil {
		return models.Task{}, false, err
	}
	if len(tasks) == 0 {
		return models.Task{}, false, models.ErrNotFound
	}
	return tasks[0], len(tasks) == 2, nil
}
