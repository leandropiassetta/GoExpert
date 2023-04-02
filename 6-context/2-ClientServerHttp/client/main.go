package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// i´m have a context that will expired in 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// independs will expired or no, i will go to cancel this context.
	defer cancel()

	// i´m will call my server http, that is, it will call this URL and if us not receive the return of this URL in 5 seconds the context will be canceled.

	// here  im prepared my request
	request, error := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if error != nil {
		panic(error)
	}

	// Do() -> this is to make my request
	result, error := http.DefaultClient.Do(request)

	if error != nil {
		panic(error)
	}

	// close my conexion
	defer result.Body.Close()

	// get the result in the body and show me in the browser with os.Stdout
	io.Copy(os.Stdout, result.Body)
}
