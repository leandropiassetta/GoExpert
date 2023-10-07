package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ZipCodeBrasil struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// i dont up a server http without a route, i need of somethint for attach to one server http, for make this i need functions and these functions need of interfaces for to implement them
	http.HandleFunc("/", SearchZipCodeHandler)

	// one line i can go up a server http
	http.ListenAndServe(":8088", nil)
}

// writer -> every response and i am want send the user, i put the data here

// reader -> every request that i will receiver i can get this data here
func SearchZipCodeHandler(writer http.ResponseWriter, reader *http.Request) {
	if reader.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// // i can manipulate a QueryString
	zipCodeParam := reader.URL.Query().Get("zip-code")

	if zipCodeParam == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	zipCode, error := SearchZipCode(zipCodeParam)
	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return in the format JSON
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// will Encoder my response(writer) the result of ZIPCODE and throw it in my response(writer) of this my requisition

	// im utilize this when i want throw this result directly and not for any variable for this is better make with json.Marshal()

	json.NewEncoder(writer).Encode(zipCode)
}

func SearchZipCode(zipCode string) (*ZipCodeBrasil, error) {
	// here i am hit in the viaCep ENDPOINT
	response, error := http.Get("https://viacep.com.br/ws/" + zipCode + "/json/")

	if error != nil {
		return nil, error
	}

	// this function will be the last to be called and will close the connection with the server http not to have resource leak
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	var zipCodeBrasil ZipCodeBrasil

	error = json.Unmarshal(body, &zipCodeBrasil)
	if error != nil {
		return nil, error
	}

	return &zipCodeBrasil, nil
}
