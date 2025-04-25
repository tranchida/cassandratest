package model

type Repository interface {
	GetHotels() ([]Hotel, error)
	GetHotel(id string) (Hotel, error)
	Close()
}
