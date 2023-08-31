package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rmit-publics/zipcode.git/internal/zipcode/infra/repository"
	"github.com/rmit-publics/zipcode.git/internal/zipcode/usecase"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/address/{zipcode}", getAddress)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func getParam(req *http.Request, param string) string {
	vars := mux.Vars(req)
	paramData := vars[param]
	return paramData
}

func getAddress(w http.ResponseWriter, req *http.Request) {
	zipcode := getParam(req, "zipcode")
	repository := repository.NewGetAddressViaCepRepository()

	input := usecase.GetAddressByZipCodeInput{Zipcode: zipcode}
	getAddressByZipCodeUseCase := usecase.NewGetAddressByZipCode(repository)
	resp, err := getAddressByZipCodeUseCase.Execute(input)
	if err != nil {
		return
	}

	output, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
