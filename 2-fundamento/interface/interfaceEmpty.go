package main

import "fmt"

// interface empyted

// Go have today in day, the generics

// ter mais liberdade nos tipos de dados que eu quero passar, por isso é bom trabalhar com interface vazia
// utilize com moderação, a vantagem do Go é trabalhar com tipagem forte.

// uso para simular um tipo de dado, bom para mocks
type x interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
