package task

import "errors"

func (s *srv) Delete(name string) error {
	if name == "" {
		return errors.New("value can't be empty")
	}
	return s.repo.Delete(name)
}
