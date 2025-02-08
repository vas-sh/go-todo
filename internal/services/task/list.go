package task

import "context"

func (s *srv) List(ctx context.Context) ([]string, error) {
	return s.repo.List(ctx)
}
