package main

import (
	"fmt"

	mathematics "packageModules/mathematics"
	portuguese "packageModules/portuguese"
)

// When created a project in Go, the first thing that we have make is created the package go mod with command "go mod init"
// the ideal is created the module with name with your github address

func main() {
	sum := mathematics.Sum(10.10, 10.10)
	loveInPortuguese, err := portuguese.TranslateLove("love")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("The result of sum is: %v \n", sum)
	fmt.Printf("The translate of love in Portuguese is? %v", loveInPortuguese)
}
