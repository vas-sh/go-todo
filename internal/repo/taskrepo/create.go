package taskrepo

import "github.com/vas-sh/todo/internal/models"

func (r *repo) Create(title string) error {
	return r.db.Create(&models.Task{
		Title:  title,
		Status: models.NewStatus,
	}).Error
}
