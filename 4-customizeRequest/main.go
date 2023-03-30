package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	// one thing that need to be clear, the request and the data of this request are distinct things, i have the data of the request and i have the client that use this data of the request to make the call

	// configurated one request here:
	c := http.Client{}

	request, error := http.NewRequest("GET", "http://google.com", nil)

	if error != nil {
		panic(error)
	}

	// add new things in my request:

	request.Header.Set("Accept", "application/json")

	// and now i have the object of the request and i have my http.Client how i make the link in this two things?
	response, error := c.Do(request)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(request.Body)

	if error != nil {
		panic(error)
	}

	println(string(body))
}
