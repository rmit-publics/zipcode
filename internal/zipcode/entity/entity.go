package entity

type ZipCodeRepository interface {
	Get(zipcode string) (*Address, error)
}

type Address struct {
	Zipcode      string
	Logradouro   string
	Neighborhood string
	State        string
	City         string
}
