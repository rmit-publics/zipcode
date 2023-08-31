package usecase

import (
	"github.com/rmit-publics/zipcode.git/internal/zipcode/entity"
)

type GetAddressByZipCodeInput struct {
	Zipcode string `json:zip_code`
}

type GetAddressByZipCodeOutPut struct {
	Zipcode      string
	Logradouro   string
	Neighborhood string
	State        string
	City         string
}

type GetAddressByZipCodeUseCase struct {
	Repository entity.ZipCodeRepository
}

func NewGetAddressByZipCode(respository entity.ZipCodeRepository) *GetAddressByZipCodeUseCase {
	return &GetAddressByZipCodeUseCase{
		Repository: respository,
	}
}

func (u *GetAddressByZipCodeUseCase) Execute(input GetAddressByZipCodeInput) (*GetAddressByZipCodeOutPut, error) {
	address, err := u.Repository.Get(input.Zipcode)
	if err != nil {
		return nil, err
	}

	return &GetAddressByZipCodeOutPut{
		Zipcode:      address.Zipcode,
		Logradouro:   address.Logradouro,
		Neighborhood: address.Neighborhood,
		State:        address.State,
		City:         address.City,
	}, nil
}
