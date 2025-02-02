package task

func (s *srv) List() ([]string, error) {
	return s.repo.List()
}
