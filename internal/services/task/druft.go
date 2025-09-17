package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) CreateTaskDruft(ctx context.Context, body models.TaskDruft) error {
	return s.repo.CreateTaskDruft(ctx, body)
}

func (s *srv) DeleteTaskDruft(ctx context.Context, userID int64) error {
	return s.repo.DeleteTaskDruft(ctx, userID)
}

func (s *srv) GetTaskDruftStatus(ctx context.Context, userID int64) (models.UserStatus, error) {
	return s.repo.GetTaskDruftStatus(ctx, userID)
}

func (s *srv) UpdateTaskDruft(ctx context.Context, body models.TaskDruft) error {
	return s.repo.UpdateTaskDruft(ctx, body)
}

func (s *srv) FindTaskDruft(ctx context.Context, userID int64) error {
	return s.repo.FindTaskDruft(ctx, userID)
}

func (s *srv) CreateFromDruft(ctx context.Context, userID int64) error {
	return s.repo.CreateFromDruft(ctx, userID)
}
