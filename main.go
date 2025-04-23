package main

import (
	"cassandratest/internal/database"
	"cassandratest/internal/server"
)

func main() {

	database, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	server := server.NewServer()
	server.DB = *database

	server.Router.Run(":8080")

}
