package databace

import (
	"database/sql"
	"log"
)

func Database() *sql.DB {
	DSN := "host=localhost user=vas password=2222 dbname=test_db port=5432 sslmode=disable TimeZone=Europe/Kiev"

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetTasksFromDB() ([]string, error) {
	db := Database()
	defer db.Close()

	rows, err := db.Query("SELECT my_task FROM task;")
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

func AddTaskToDB(task string) {
	db := Database()
	defer db.Close()

	f := `INSERT INTO task (my_task) VALUES ($1)`
	if task == "" {
		log.Println("Value can't be empty")
		return
	}

	_, err := db.Exec(f, task)
	if err != nil {
		log.Printf("Error inserting value to db: %s", err)
		return
	}
	log.Println("Task secsesfuly added")
}
