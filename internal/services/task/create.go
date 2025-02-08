package task

import (
	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Create(title string) error {
	if title == "" {
		return models.ErrValueEmpty
	}
	return s.repo.Create(title)
}
