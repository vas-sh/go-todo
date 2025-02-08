package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Create(ctx context.Context, title string) error {
	if title == "" {
		return models.ErrValueEmpty
	}
	return s.repo.Create(ctx, title)
}
