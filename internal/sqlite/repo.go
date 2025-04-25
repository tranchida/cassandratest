package sqlite

import (
	"cassandratest/internal/model"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteRepo struct {
	db *sql.DB
}

func NewSqliteRepo() (model.Repository, error) {
	db, err := sql.Open("sqlite3", "./hotels.db")
	if err != nil {
		return nil, err
	}

	err = initDatabase(db)
	if err != nil {
		return nil, err
	}

	return &SqliteRepo{db: db}, nil
}

func (repo *SqliteRepo) GetHotels() ([]model.Hotel, error) {
	rows, err := repo.db.Query("SELECT id, name, phone, street, city, state_or_province, postal_code, country FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hotels := []model.Hotel{}
	for rows.Next() {
		var h model.Hotel
		var street, city, state, postal, country string
		if err := rows.Scan(&h.Id, &h.Name, &h.Phone, &street, &city, &state, &postal, &country); err != nil {
			return nil, err
		}
		h.Address = model.Address{
			Street:          street,
			City:            city,
			StateOrProvince: state,
			PostalCode:      postal,
			Country:         country,
		}
		hotels = append(hotels, h)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return hotels, nil
}

func (repo *SqliteRepo) GetHotel(id string) (model.Hotel, error) {

	hotel := model.Hotel{}

	err := repo.db.QueryRow("SELECT id, name, phone, street, city, state_or_province, postal_code, country FROM hotels WHERE id = ?", id).Scan(&hotel.Id, &hotel.Name, &hotel.Phone, &hotel.Address.Street, &hotel.Address.City, &hotel.Address.StateOrProvince, &hotel.Address.PostalCode, &hotel.Address.Country)
	if err != nil {
		return hotel, err
	}

	return hotel, nil
}

func (repo *SqliteRepo) Close() {
	repo.db.Close()
}

func initDatabase(db *sql.DB) error {

	// check if hotel table exist
	var tableExist bool
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='hotels'`).Scan(&tableExist)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if !tableExist {
		_, err = db.Exec(`CREATE TABLE hotels (
			id text PRIMARY KEY,
			name text,
			phone text,
			street text,
			city text,
			state_or_province text,
			postal_code text,
			country text
		)`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`INSERT INTO hotels (id, name, phone, street, city, state_or_province, postal_code, country) 
		VALUES ('mirador', 'Hotel Mirador','1234', 'Rte des Alpes 10', 'Montreux', '', '1820', 'Switzerland'),
		('rothay', 'Rothay Manor Hotel','+44 1539 445759', 'Rothay Bridge', 'Ambleside', 'Cumbria', 'LA22 0EH', 'United Kingdom')`)
		if err != nil {
			return err
		}
	}

	return nil
}
