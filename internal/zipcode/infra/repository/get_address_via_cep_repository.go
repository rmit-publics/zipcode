package repository

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/rmit-publics/zipcode.git/internal/zipcode/entity"
)

type AddressViaCep struct {
	Zipcode      string `json:"cep"`
	Logradouro   string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
}

type GetAddressViaCepRepository struct {
	url string
}

func NewGetAddressViaCepRepository() *GetAddressViaCepRepository {
	return &GetAddressViaCepRepository{
		url: "https://viacep.com.br/ws/",
	}
}

func (u *GetAddressViaCepRepository) Get(zipcode string) (*entity.Address, error) {
	request := u.url + zipcode + "/json"
	resp, err := http.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var address AddressViaCep

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(bodyBytes, &address)
	}

	var addressResponse entity.Address
	addressResponse.Zipcode = address.Zipcode
	addressResponse.Logradouro = address.Logradouro
	addressResponse.Neighborhood = address.Neighborhood
	addressResponse.State = address.State
	addressResponse.City = address.City

	return &addressResponse, nil
}
