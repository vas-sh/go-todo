package taskrepo

func (r *repo) Create(name string) error {
	f := `INSERT INTO task (my_task) VALUES ($1)`
	_, err := r.db.Exec(f, name)
	return err
}
