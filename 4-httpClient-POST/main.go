package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}

	// POST -> will need of body of my request
	// my body is a io.reader this read a slice of bytes one slice of bytes i will can transformed in a string, when i need send data for comunicate in two system utilized of http protocol, i sendo this datas in bytes(format JSON)

	jsonFormat := bytes.NewBuffer([]byte(`{"name":"leandro"}`))

	response, error := c.Post("http://google.com", "application/json", jsonFormat)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	// CopyBuffer -> take the data and i m choice where i want to throw this data and where i will go copy this data

	// os.Stdout -> where i will to throw this data
	// response.Body -> where i will copy this data
	io.CopyBuffer(os.Stdout, response.Body, nil)
}
