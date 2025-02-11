package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Create(ctx context.Context, title, description string) error {
	if title == "" {
		return models.ErrValueEmpty
	}
	return s.repo.Create(ctx, title, description)
}
