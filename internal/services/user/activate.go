package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Activate(ctx context.Context, id uuid.UUID) (models.User, error) {
	userActivation, err := s.repo.Activation(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	if time.Since(userActivation.Date).Hours() > 2 {
		return models.User{}, models.ErrAlreadyExpired
	}
	if userActivation.Activated {
		return models.User{}, models.ErrAlreadyActivated
	}
	user := userActivation.User
	return user, s.repo.Activate(ctx, userActivation)
}
