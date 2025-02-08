package taskrepo

import "github.com/vas-sh/todo/internal/models"

func (r *repo) Remove(name string) error {
	return r.db.Model(models.Task{}).Delete(models.Task{}, "title = ?", name).Error
}
