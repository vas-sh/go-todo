package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) CreateTaskDruft(ctx context.Context, body models.TaskDruft) error {
	return r.db.WithContext(ctx).Create(&body).Error
}

func (r *repo) DeleteTaskDruft(ctx context.Context, userID int64) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(models.TaskDruft{}).Error
}

func (r *repo) UpdateTaskDruft(ctx context.Context, body models.TaskDruft) error {
	return r.db.WithContext(ctx).Where("user_id = ?", body.UserID).Updates(&body).Error
}

func (r *repo) FindTaskDruft(ctx context.Context, userID int64) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).First(&models.TaskDruft{}).Error
}

func (r *repo) GetTaskDruftStatus(ctx context.Context, userID int64) (models.UserStatus, error) {
	var status models.UserStatus
	err := r.db.WithContext(ctx).Model(models.TaskDruft{}).
		Select("user_status").
		Where("user_id = ?", userID).
		Scan(&status).Error
	return status, err
}
