package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Activate(ctx context.Context, id uuid.UUID) error {
	userActivation, err := s.repo.Activation(ctx, id)
	if err != nil {
		return err
	}
	if time.Since(userActivation.Date).Hours() > 2 {
		return models.ErrAlreadyExpired
	}
	if userActivation.Activated {
		return models.ErrAlreadyActivated
	}
	return s.repo.Activate(ctx, userActivation)
}
