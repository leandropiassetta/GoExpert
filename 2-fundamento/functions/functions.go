package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(sumIsPair(1, 2))

	number, isPair, err := sumIsPair(2, 3)
	if err != nil {
		fmt.Println(err)
	}

	if isPair {
		fmt.Println(number)
	}

	pairNumber := number + number

	fmt.Println(pairNumber)
}

func sumIsPair(a, b int) (int, bool, error) {
	if (a + b) < 100 {
		return (a + b), false, errors.New("this number is not allowed")
	}

	if (a+b)%2 != 0 {
		return (a + b), false, nil
	}
	return (a + b), true, nil
}
