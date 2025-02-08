package task

import "context"

type repoer interface {
	Create(ctx context.Context, title string) error
	Remove(ctx context.Context, title string) error
	List(ctx context.Context) ([]string, error)
}

type srv struct {
	repo repoer
}

func New(repo repoer) *srv {
	return &srv{repo: repo}
}
