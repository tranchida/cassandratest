package main

import (
	"cassandratest/internal/database"
	"cassandratest/internal/server"
	_ "cassandratest/docs"
)

// @title CassandraTest API
// @version 1.0
// @description This is a sample server for CassandraTest.
// @host localhost:8080
// @BasePath /

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
