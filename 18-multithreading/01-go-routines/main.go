package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running \n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")

	// Thread 4
	// anonymous function
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running \n", i, "anonymous function")
			time.Sleep(1 * time.Second)
		}
	}()

	// Nothing here ...

	// My programming main comes here and it ends immediately after that because there is nothing to do here. So, the program ends and all the threads are killed.

	// So, we need to wait for the threads to finish their tasks. We can do that by using time.Sleep() function.
	time.Sleep(10 * time.Second)
}
