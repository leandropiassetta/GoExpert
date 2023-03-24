package main

import "fmt"

// Quando você passa uma interface vazia para uma função, o Go vai querer saber qual é o tipo daquilo que você colocou, e você pode forçar aquele tipo ser "transformado", para saber que Go saiba exatamente daquilo que o Go esta falando.

func main() {
	var minhaVar interface{} = "Leandro de Freitas"

	// estou afirmando pro meu sistema que essa variável é uma string.
	println(minhaVar.(string))

	// fazendo uma verificação para ve se minha conversão deu certo ou não.
	res, ok := minhaVar.(int)

	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)

	res2 := minhaVar.(int)

	fmt.Printf("O valor de res é %v", res2)
}
