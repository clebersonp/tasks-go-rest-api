package models

import "github.com/clebersonp/tasks-go-rest-api/db"

// Get returns one task by it id
func Get(id int64) (task Task, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM tasks WHERE id=$1`, id)
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.Done)

	return
}
