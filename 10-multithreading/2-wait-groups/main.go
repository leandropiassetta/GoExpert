package main

import (
	"fmt"
	"sync"
	"time"
)

// A WaitGroup in Go is a control structure used to synchronize the execution of goroutines(lightweight threads) in a concurrent program. It is particularly useful when you want to wait for a group of goroutines to finish their execution before proceeding with another part of the program.

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running \n", i, name)
		time.Sleep(1 * time.Second)

		// Done(): Within each goroutine, when it finishes its task, it calls the Done() method on the WaitGroup. This decrements the internal counter in WaitGroup by 1.
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}

	// Add(): You can add goroutines to the WaitGroup using the Add() method. Each call to Add() increments an internal counter in WaitGroup, allowing you to indicate how many go routines you expect to be completed.

	// Add 25 goroutines to the WaitGroup
	waitGroup.Add(25)
	// Thread 2
	go task("A", &waitGroup)
	// Thread 3
	go task("B", &waitGroup)

	// Thread 4
	// anonymous function
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running \n", i, "anonymous function")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	// Wait(): The Wait() method blocks the execution of the main program or another go routine until the WaitGroup counter reaches zero. This means that the program will wait until all the goroutines added with Add() method have called Done() method.
	waitGroup.Wait()
}
