package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

// in go haver a package called CONTEXT, what the context does?
// its a package that allows us to pass information on it for diversous calls to our system and also the option for that this context to be cancelled and when this context is cancelled the application stop on time for dont spend too much time
func main() {
	// by default we use this variable "ctx "when are we going utilize context

	// new context
	ctx := context.Background()

	// if one application to run using context, if pass one second, this context to be cancelled
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	// other way of cancel this context
	defer cancel()

	// this request only execute if my rule of my context pass here
	request, error := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)

	if error != nil {
		panic(error)
	}

	response, error := http.DefaultClient.Do(request)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)

	println(string(body))
}
