package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Create(ctx context.Context, res *models.Task) error {
	err := r.db.WithContext(ctx).Create(res).Error
	return err
}

func (r *repo) CreateFromDruft(ctx context.Context, userID int64) error {
	var task models.TaskDruft
	err := r.db.WithContext(ctx).Model(models.TaskDruft{}).Where("user_id = ?", userID).First(&task).Error
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(models.Task{}).Create(map[string]any{
		"title":         task.Title,
		"description":   task.Description,
		"status":        task.Status,
		"user_id":       userID,
		"estimate_time": task.EstimateTime,
	}).Error
}
