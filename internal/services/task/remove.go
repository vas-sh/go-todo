package task

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Remove(ctx context.Context, title string) error {
	if title == "" {
		return models.ErrValueEmpty
	}
	return s.repo.Remove(ctx, title)
}
