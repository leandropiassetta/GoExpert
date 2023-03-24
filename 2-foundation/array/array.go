package main

import "fmt"

const a = "hello world"

type ID int

var (
	b bool   = true
	c int    = 10
	d string = "Leandro"
	f ID     = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", f)

	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	// ultimo elemento dinamico

	fmt.Println(meuArray[len(meuArray)-1])

	// percorrendo Array

	for i, v := range meuArray {
		fmt.Printf("o valor do indice %d é %d\n", i, v)
	}
}
