package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64

// concurrent issues in practice

// this server for each request it create a new thread, and how this server create multiple threads as it generates many requests at the same time, so concurrent can happen and there can be a problem where at the exact moment two threads are reading the same value of number and incrementing it, so the value of number will not be correct.

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++

		time.Sleep(300 * time.Millisecond)
		w.Write([]byte("You is the visitor number " + fmt.Sprint(number) + "\n"))
	})

	http.ListenAndServe(":3000", nil)
}
