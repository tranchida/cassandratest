package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
)

type Database struct {
	Session *gocql.Session
	Ctx     context.Context
}

func NewDatabase() (*Database, error) {

	hostname := os.Getenv("CASSANDRA_HOST")
	if hostname == "" {
		hostname = "localhost"
	}
	port := os.Getenv("CASSANDRA_PORT")
	if port == "" {
		port = "9042"
	}

	cluster := gocql.NewCluster(fmt.Sprintf("%s:%s", hostname, port))
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println("Error creating Cassandra session:", err)
		return nil, err
	}

	ctx := context.Background()

	return &Database{
		Session: session,
		Ctx:     ctx,
	}, nil
}

func (database *Database) Close() {
	if database.Session != nil {
		database.Session.Close()
	}
}
