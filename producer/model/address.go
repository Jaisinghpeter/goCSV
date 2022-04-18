package model

type Address struct {
    Country string `json:"country"`
	City string `json:"city"`
	Place string `json:"place"`
	Zip string `json:"zipcode"`
}

func NewAddress( country, city, place, zip string) *Address {
    address := new(Address)
    address.City = city
	address.Country = country
	address.Place = place
	address.Zip = zip
    return address
}

const(
	Country = 30
	City = 31
	Place = 29
	Zip = 33
)