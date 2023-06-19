package models

import "github.com/clebersonp/tasks-go-rest-api/db"

// Delete deletes a task by the given id and returns the number of rows affected or the error
func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM tasks WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}