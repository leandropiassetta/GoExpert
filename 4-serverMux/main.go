package main

// SERVERMUX
// Multiplexer(MUX) -> a component that attach routes in him

// It is a device that selects information from two or more data sources into a single channel.

// I CAN CREATE MY OWN SERVERMUX IN GO, for i have more control about my application
import (
	"net/http"
)

func main() {
	// this function NewServeMux its responsable for attach the handlers of my App

	mux := http.NewServeMux()

	// anonymous function
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})

	http.ListenAndServe(":8089", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Leandro!!!"))
	})

	// the URL are diferents because they are distinct mux
	http.ListenAndServe(":8090", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

// always that i am will create a handler, this handler implement an interface of Handler and this interface force us to have a method SERVERHTTTP
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
