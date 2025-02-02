package task

import "errors"

func (s *srv) Create(name string) error {
	if name == "" {
		return errors.New("value can't be empty")
	}
	return s.repo.Create(name)
}
