package main

import "fmt"

func main() {
	// fmt.Print(sum(1, 2, 3, 4, 5, 6, 6, 7, 7, 8, 109))
	fmt.Println(IsPair(2, 10, 20, 33, 35))
}

// exemplos que podem usar esse tipo de função variádicas, passar alguns parametros de configuração pro seu banco de dados, eles podem ser 1 como 10 parametros, não sabe a quantidade coloca os ...
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func IsPair(numbers ...int) []int {
	sum := sum(numbers...)
	control := 0
	num1 := 0
	num2 := 0

	pairs := make([]int, 0, sum/2)
	// simuling while
	for control < sum {

		if (num1+num2)%2 == 0 {
			pair := num1 + num2
			pairs = append(pairs, pair)
			num1 += 1
			control += 1

		}

		num2 += 1
		control += 1
	}

	fmt.Println(len(pairs), cap(pairs))
	return pairs
}
