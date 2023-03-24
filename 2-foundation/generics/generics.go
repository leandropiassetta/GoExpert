package main

// the generics helps when we have this kind of situation, where we have two similar functions that only change of type of them

func sumSalariesInt(mapsalaries map[string]int) int {
	var sum int

	for _, value := range mapsalaries {
		sum += value
	}

	return sum
}

func sumSalariesFloat(mapsalaries map[string]float64) float64 {
	var sum float64

	for _, value := range mapsalaries {
		sum += value
	}

	return sum
}

// GENERICS:

// when this happen, that situation could be resolved whith generics:

func SumGeneric[T int | float64](mapsalaries map[string]T) T {
	var sum T

	for _, value := range mapsalaries {
		sum += value
	}

	return sum
}

// CONSTRAINTS:

// Are types specific that i create and i could replace into of paramater of my function generics

type Number interface {
	~int | ~float64
}

// the tilde accent "~", we will use when we have a situation where can we have other type int

type MyNumber int

func SumWithConstraints[T Number](mapsalaries map[string]T) T {
	var sum T

	for _, value := range mapsalaries {
		sum += value
	}

	return sum
}

// comparable is a constraint will accept the T's as long as these T´s will can comparables
// comparable -> will bring values that can b compared
// comparable only ""=="" and not "< , >"
func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	salaries := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	salaries2 := map[string]MyNumber{"Wesley": 2000, "João": 3000, "Maria": 4000}

	salariesFloat := map[string]float64{"Wesley": 100.20, "João": 200.30, "Maria": 300.40}

	// fmt.Printf("The sum of all employees salaries is: %d", sumSalariesInt(salaries))
	// fmt.Printf("The sum of all employees salaries is: %f", sumSalariesFloat(salariesFloat))

	println(SumGeneric(salaries))
	println(SumWithConstraints(salariesFloat))
	println(SumWithConstraints(salaries2))

	println(Compare(10, 10))
}
