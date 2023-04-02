package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// the we request also received one Context() inside her, and this context can many information in her, with "correlation id" for example
	ctx := r.Context()

	log.Println("initialized request")

	defer log.Println("completed request")

	select {
	case <-time.After(5 * time.Second):
		// this log -> print in command line stdout
		log.Println("request successfully processed")
		// Print in the browser
		w.Write([]byte("request successfully processed"))
	case <-ctx.Done():
		// this log -> print in command line stdout
		log.Println("request canceled for client")
	}
}

// CONTEXT -> prevents us from doing unnecessary processing as time goes by
