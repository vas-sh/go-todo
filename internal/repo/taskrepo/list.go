package taskrepo

func (r *repo) List() ([]string, error) {
	rows, err := r.db.Query("SELECT my_task FROM task;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []string
	for rows.Next() {
		var task string
		err := rows.Scan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}
