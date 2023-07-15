package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	timeout := 300 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8082", nil)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout exceeded. Execution time was insufficient.")
		}
		panic(err)
	}

	// make the request
	result, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer result.Body.Close()

	// print the result to the terminal window (stdout)

	io.Copy(os.Stdout, result.Body)
}
