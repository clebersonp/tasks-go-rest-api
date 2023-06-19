// Package that provides database connection
package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/clebersonp/tasks-go-rest-api/configs"
	_ "github.com/lib/pq" // import the postgresql driver
)

// OpenConnection opens database connection based on configuration
func OpenConnection() (*sql.DB, error) {
	conf := configs.DB()
	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", urlConnection)
	if err != nil {
		log.Fatalln("Could not open postgresql database connection:", err)
	}

	err = conn.Ping()

	return conn, err
}