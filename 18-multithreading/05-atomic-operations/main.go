package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64

// ATOMIC OPERATIONS: Are operations that are executed in a single step, without the possibility of interruption.

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// number++

		// In this package make internal operations to avoid the interruption of the goroutines.
		atomic.AddUint64(&number, 1)

		time.Sleep(300 * time.Millisecond)
		w.Write([]byte("You is the visitor number " + fmt.Sprint(number) + "\n"))
	})

	http.ListenAndServe(":3000", nil)
}
