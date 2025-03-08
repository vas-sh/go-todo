package userrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) GetByEmail(ctx context.Context, email string) (user models.User, err error) {
	err = r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return
}
