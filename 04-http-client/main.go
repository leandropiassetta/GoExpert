package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

// When i need created a fast system the first thing that i need to think is how do i leave this system more perfomatic possible and a way to do this is manage to establish the limit the external called.

// For example if a called for an EXTERNAL API take longer than 10 seconds for return a result so its important estabilish limit, sayed for us sytem that this API take longer more than 2 second sayed her return a error.

func main() {
	// i can pass many parameters for http.Client, one this parameters is the timeOut -> maximum time this application can take to return a response
	c := http.Client{Timeout: time.Millisecond}

	response, error := c.Get("http://google.com")

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	println(string(body))
}
