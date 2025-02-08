package task

import (
	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Remove(title string) error {
	if title == "" {
		return models.ErrValueEmpty
	}
	return s.repo.Remove(title)
}
