package taskrepo

func (r *repo) Delete(name string) error {
	f := "DELETE FROM task WHERE my_task = ($1)"
	_, err := r.db.Exec(f, name)
	return err
}
