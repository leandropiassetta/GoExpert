package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // In this moment the channel is empty.

	// Thread 2
	go func() {
		channel <- "Hello World!" // In this moment the channel is full.
	}()

	// Thread 1

	msg := <-channel // In this moment the channel is empties.

	fmt.Println(msg)
}
