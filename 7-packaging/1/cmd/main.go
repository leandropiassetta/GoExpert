package main

import (
	"fmt"

	"github.com/leandropiassetta/goexpert/7-packaging/1/math"
)

// cmd -> is a folder that contains the main.go file that is the entry point of the application and is the file that is compiled and executed when we run the command go run main.go

// cmd -> contains archives of execution of the application, archives that are executed when we run the command go run main.go

func main() {
	math := math.NewMath(1, 2, "mathematic")

	fmt.Println(math)
	fmt.Println(math.Class())
	fmt.Println(math.Sum())
}
