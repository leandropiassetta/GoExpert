package main

import (
	"crypto/rand"
	"math/big"
)

func retryFor() {
	// many times we need to retry something until it works

	aleatoryNumber, _ := rand.Int(rand.Reader, big.NewInt(10))

	manyTimes := 1000

	for i := 0; i < manyTimes; i++ {
		if aleatoryNumber.Int64() != int64(i) {
			manyTimes = manyTimes / 10
		} else {
			println("found", i)
		}

		if manyTimes == 1 {
			println("It worked!")
			break
		}
	}
}

func main() {
	retryFor()

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
