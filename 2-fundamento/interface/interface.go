package main

import (
	"fmt"
)

// Significa automaticamente que qualquer struct que tiver um método desativar estaria implementando a interface pessoa.
type Pessoa interface {
	Desativar()
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (empresa Empresa) Desativar() {
}

type Empresa struct {
	Nome string
}

func (client Client) Desativar() {
	client.Ativo = false
	fmt.Printf("O cliente %s foi desativado", client.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	// leandro := Client{
	// 	Nome:  "Leandro",
	// 	Idade: 36,
	// 	Ativo: true,
	// }

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)
}
