package main

func main() {
	// In Go we have only for and for range

	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []string{"one", "two", "three"}

	for k, v := range numbers {
		println(k, v)
	}

	// simulating a while
	i := 0

	for i < 10 {
		println(i)
		i++
	}

	// infinite looping:

	for {
		println("Hello, World!!")
	}

	// When do we use it??

	// Let's imagine you want to consume a message from a queue, use something that will never stop.
}
