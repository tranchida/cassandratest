package cassandra

import (
	"cassandratest/internal/model"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
)

type CassandraRepo struct {
	Session *gocql.Session
	Ctx     context.Context
}

func NewCassandraRepo() (model.Repository, error) {

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

	return &CassandraRepo{
		Session: session,
		Ctx:     ctx,
	}, nil
}

func (repo *CassandraRepo) Close() {
	if repo.Session != nil {
		repo.Session.Close()
	}
}

func (repo CassandraRepo) GetHotels() ([]model.Hotel, error) {

	iter := repo.Session.Query("SELECT id, address, name, phone FROM hotel.hotels").WithContext(repo.Ctx).Iter()
	defer iter.Close()

	var hotels []model.Hotel
	for {
		hotel := model.Hotel{}
		if !iter.Scan(&hotel.Id, &hotel.Address, &hotel.Name, &hotel.Phone) {
			break
		}
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}

func (repo CassandraRepo) GetHotel(id string) (model.Hotel, error) {

	hotel := model.Hotel{}

	err := repo.Session.Query("SELECT id, address, name, phone FROM hotel.hotels WHERE id = ?", id).WithContext(repo.Ctx).Scan(&hotel.Id, &hotel.Address, &hotel.Name, &hotel.Phone)
	if err != nil {
		return hotel, err
	}

	return hotel, nil
}
