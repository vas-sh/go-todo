package userrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Create(ctx context.Context, user *models.User) error {
	var users []models.User
	err := r.db.WithContext(ctx).Where("email = ?", user.Email).Find(&users).Error
	if err != nil {
		return err
	}
	if len(users) > 0 {
		if users[0].Activated {
			return models.ErrAlreadyExist
		}
		*user = users[0]
		return nil
	}
	return r.db.WithContext(ctx).Create(user).Error
}
