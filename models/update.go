package models

import "github.com/clebersonp/tasks-go-rest-api/db"

// Update updates the given task in the database and returns the number of rows affected or the error
func Update(id int64, task Task) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `UPDATE tasks SET title=$2, description=$3, done=$4 WHERE id=$1`
	result, err := conn.Exec(sql, id, task.Title, task.Description, task.Done)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
