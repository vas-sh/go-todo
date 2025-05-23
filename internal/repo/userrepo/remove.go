package userrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) Remove(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Model(models.User{}).Delete(models.User{}, "id = ?", id).Error
}
