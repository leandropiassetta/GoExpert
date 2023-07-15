package main

import (
	"fmt"
)

// DEFER

// it's like i want to delay something

// is very common, for example, i dont remember closed a connection

func main() {
	// req, err := http.Get("https://www.google.com")
	// if err != nil {
	// 	panic(err)
	// }

	// // DEFER -> its a statement that will delay
	// defer req.Body.Close()

	// result, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("First line")
	defer fmt.Println("Second line")
	fmt.Println("Third line")
	fmt.Println("Fourthline")
}
