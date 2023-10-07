package main

// don't execute this application because him went published as a library in github.com repository, my go programm not found this package in my local machine

// for publish this package in my local machine i need execute this command:
// go mod edit -replace github.com/leandropiassetta/goexpert/7-Packaging/3/math=../math

// for remove this package in my local machine i need execute this command:
// go mod edit -dropreplace github.com/leandropiassetta/goexpert/7-Packaging/3/math

import "github.com/leandropiassetta/goexpert/7-Packaging/3/math"

func main() {
	math := math.NewMath(1, 2, "mathematic")

	println(math.Sum())
}
