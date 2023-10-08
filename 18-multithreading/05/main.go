package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++

		time.Sleep(300 * time.Millisecond)
		w.Write([]byte("You is the visitor number " + fmt.Sprint(number) + "\n"))
	})

	http.ListenAndServe(":3000", nil)
}
