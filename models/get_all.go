package models

import (
	"log"

	"github.com/clebersonp/tasks-go-rest-api/db"
)

// GetAll returns all tasks in database
func GetAll() (tasks []Task, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM tasks`)
	if err != nil {
		return
	}

	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done)
		if err != nil {
			log.Println("Something bad happened while trying to get some task:", err)
			continue // try the next
		}
		tasks = append(tasks, task)
	}

	return
}
