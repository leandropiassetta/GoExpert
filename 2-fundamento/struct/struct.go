package main

import (
	"fmt"
)

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Composicion of structs
// A Struct have can have a method
type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func IndexOf(arr []Client, candidate Client) int {
	for index, c := range arr {
		if c == candidate {
			return index
		}
	}
	return -1
}

func removeInactiveMember(client Client, membersClient []Client) []Client {
	activeMembers := membersClient
	indexClient := IndexOf(membersClient, client)

	if !client.Ativo {
		activeMembers = append(activeMembers[:indexClient], activeMembers[indexClient+1:]...)
	}

	return activeMembers
}

func (client *Client) Desativar() {
	client.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n\n", client.Nome)
}

func main() {
	leandro := Client{
		Nome:  "Leandro",
		Idade: 36,
		Ativo: false,
		Endereco: Endereco{
			Logradouro: "Rua Labareda do oeste",
			Numero:     335,
			Cidade:     "Florianopolis",
			Estado:     "SC",
		},
	}

	clienteLeandro := &leandro

	joao := Client{
		Nome:  "João",
		Idade: 45,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua Radiola Amadora",
			Numero:     326,
			Cidade:     "Florianopolis",
			Estado:     "SC",
		},
	}

	donaDelia := Client{
		Nome:  "Adélia",
		Idade: 60,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua das Avenidas",
			Numero:     326,
			Cidade:     "Ararangua",
			Estado:     "SC",
		},
	}

	joao.Desativar()
	clienteLeandro.Ativo = true
	clienteLeandro.Cidade = "Criciuma"

	clubMember := []Client{
		leandro,
		joao,
		donaDelia,
	}

	fmt.Println(removeInactiveMember(joao, clubMember))

	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", leandro.Nome, leandro.Idade, leandro.Ativo)
}
