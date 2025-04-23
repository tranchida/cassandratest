package database

type Hotel struct {
	Id      string   `cql:"id" json:"id"`
	Address Address  `cql:"address" json:"address"`
	Name    string   `cql:"name" json:"name"`
	Phone   string   `cql:"phone" json:"phone"`
	Pois    []string `cql:"pois" json:"pois"`
}

type Address struct {
	Street          string `cql:"street" json:"street"`
	City            string `cql:"city" json:"city"`
	StateOrProvince string `cql:"state_or_province" json:"state_or_province"`
	PostalCode      string `cql:"postal_code" `
	Country         string `cql:"country" json:"country"`
}

func (database Database) GetHotels() ([]Hotel, error) {

	iter := database.Session.Query("SELECT * FROM hotel.hotels").WithContext(database.Ctx).Iter()
	defer iter.Close()

	var hotels []Hotel
	for {
		hotel := Hotel{}
		if !iter.Scan(&hotel.Id, &hotel.Address, &hotel.Name, &hotel.Phone, &hotel.Pois) {
			break
		}
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}
