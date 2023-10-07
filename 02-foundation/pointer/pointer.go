package main

import "fmt"

func main() {
	// Memória -> Endereço -> Valor
	// Variável -> Ponteiro que tem um endereço na memória -> Valor

	a := 10 // abre uma caixa na memoria guardando o valor de A, e ela tem o endereço.
	println("A-1", a)
	// Toda vez que alguem quiser saber o valor de A, tem que ir na caixinha e abrir ela pra ve o que tem dentro.

	fmt.Println(&a) // Endereço = 0xc0000ac008
	var ponteiro *int = &a
	*ponteiro = 20
	println(ponteiro)
	println("A-2", a)

	b := &a
	*b = 30
	// desreferenciando o valor da memória
	println("B-1", *b)
}
