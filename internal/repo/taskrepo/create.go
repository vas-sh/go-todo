package taskrepo

import "github.com/vas-sh/todo/internal/models"

func (r *repo) Create(name string) error {
	return r.db.Create(&models.Task{
		Title:  name,
		Status: models.NewStatus,
	}).Error
}
