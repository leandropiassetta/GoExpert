package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// codes of google
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	// req.Body = this is a string of data
	// i need of a reader
	result, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

	// important for no leak datas
	req.Body.Close()
}
