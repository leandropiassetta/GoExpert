package main

import (
	"log"
	"net/http"
)

func main() {
	// How could i create a file server, serve images, css, htmls?

	// for this we have a fileserver

	fileserver := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()

	mux.Handle("/", fileserver)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from blog"))
	})

	log.Fatal(http.ListenAndServe(":8088", mux))
}
