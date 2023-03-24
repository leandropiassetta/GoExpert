package main

import "fmt"

func main() {
	salarys := map[string]int{"Leandro": 50000, "João": 1000, "Maria": 20.000}

	fmt.Println(salarys["Leandro"])

	delete(salarys, "Leandro")
	salarys["Lean"] = 100000

	fmt.Println(salarys["Lean"])

	// salary := make(map[string]int)
	salary1 := map[string]int{}

	salary1["Leandro"] = 200000

	for name, salary := range salarys {
		fmt.Printf(" The salary of %s is %d\n", name, salary)
	}
}
