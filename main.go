package main

import (
	"log"

	"github.com/clebersonp/tasks-go-rest-api/configs"
)

func main() {

	// config the log formater
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	if err := configs.Load(); err != nil {
		log.Fatal(err)
	}

	dbConfig := configs.DB()
	log.Printf("DB Config: Host: %v, Port: %v, User: %v, Password: %v, Database: %v\n",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database)

	serverPort := configs.ServerPort()
	log.Printf("Server Port: %v\n", serverPort)
}
