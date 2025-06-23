package userrepo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vas-sh/todo/internal/models"
	"gorm.io/gorm"
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

func (r *repo) Activation(ctx context.Context, id uuid.UUID) (*models.UserActivation, error) {
	var res models.UserActivation
	err := r.db.WithContext(ctx).Preload("User").First(&res, id).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *repo) Activate(ctx context.Context, userActivation *models.UserActivation) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.
			Model(userActivation).
			Where("id = ?", userActivation.ID).
			Update("activated", true).Error
		if err != nil {
			return err
		}
		return tx.
			Model(models.User{}).
			Where("id = ?", userActivation.UserID).
			Update("activated", true).Error
	})
}
