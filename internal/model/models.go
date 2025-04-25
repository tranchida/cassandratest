package model

type Hotel struct {
	Id      string   `cql:"id" json:"id"`
	Address Address  `cql:"address" json:"address"`
	Name    string   `cql:"name" json:"name"`
	Phone   string   `cql:"phone" json:"phone"`
}

type Address struct {
	Street          string `cql:"street" json:"street"`
	City            string `cql:"city" json:"city"`
	StateOrProvince string `cql:"state_or_province" json:"state_or_province"`
	PostalCode      string `cql:"postal_code" `
	Country         string `cql:"country" json:"country"`
}
