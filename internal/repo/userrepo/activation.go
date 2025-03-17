package userrepo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) CreateActivation(ctx context.Context, userID int64) (uuid.UUID, error) {
	var activations []models.UserActivation
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND date >= now() - interval '1 hour' AND activated = false", userID).
		Order("date DESC").
		Limit(1).
		Find(&activations).Error
	if err != nil {
		return uuid.UUID{}, err
	}
	if len(activations) > 0 {
		return activations[0].ID, nil
	}
	activation := models.UserActivation{
		ID:     uuid.New(),
		UserID: userID,
		Date:   time.Now(),
	}
	err = r.db.WithContext(ctx).Create(&activation).Error
	return activation.ID, err
}
