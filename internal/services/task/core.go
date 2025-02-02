package task

type repoer interface {
	Create(name string) error
	Delete(name string) error
	List() ([]string, error)
}

type srv struct {
	repo repoer
}

func New(repo repoer) *srv {
	return &srv{repo: repo}
}
