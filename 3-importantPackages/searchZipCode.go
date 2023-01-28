package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// in this programm i utilized a url of viacep.com.br
// http://viacep.com.br/ws/01001000/json/

// You will go digite the zipCode and my system will show for you the name of street, city , country, etc..

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
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when making request: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		}

		var data ZipCodeBrasil

		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing response: %v\n", err)
		}

		if data.Localidade == "" {
			fmt.Fprint(os.Stderr, "This Zip Code don't exist")
			break
		}

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("ZipCode: %s, Locality: %s, State: %s", data.Cep, data.Localidade, data.Uf))
		fmt.Println("File created successfully!")
		fmt.Println("City:", data.Localidade)
	}
}
