package user

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Remove(ctx context.Context, id int64) error {
	if id <= 0 {
		return models.ErrValueEmpty
	}
	return s.repo.Remove(ctx, id)
}
