package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64

func main() {
	mutex := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// mutex Lock is used to block the access to the variable number, in this moment only one goroutine can access to the variable number.
		mutex.Lock()
		number++

		// mutex Unlock is used to unlock the access to the variable number, in this moment the goroutine that was blocked can access to the variable number.
		mutex.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte("You is the visitor number " + fmt.Sprint(number) + "\n"))
	})

	http.ListenAndServe(":3000", nil)
}
