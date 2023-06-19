package models

import (
	"github.com/clebersonp/tasks-go-rest-api/db"
)

// Insert inserts into tasks table the task
func Insert(task Task) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO tasks (title, description, done) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(sql, task.Title, task.Description, task.Done).Scan(&id)

	return
}
