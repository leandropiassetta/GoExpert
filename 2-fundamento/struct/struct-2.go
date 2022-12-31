package main

import "fmt"

// its commom will make alterations in the structs
type Client struct {
	name string
}
type BankAccount struct {
	balance int
}

// Constructor
// What does it mean?
func NewAccount() *BankAccount {
	// Create a function that return the address of the memory of the account that is will return
	return &BankAccount{balance: 0}
}

// Why if make this??

// Because now means that anywhere where will pass this bankaccount, and the program will make a alteration this will be reflect in anywhere in the program, it will not be another copy

func (client *Client) walk() {
	client.name = "Leandro Piassetta"
	fmt.Printf("The client %v walked\n", client.name)
}

// To simulate is ideal that don't alter the value of client bank balance but the copy of value.
// For that doesn't happen use pointer.
func (client BankAccount) simulate(value int) int {
	client.balance += value
	println(client.balance)
	return client.balance
}

func main() {
	clienBank1 := BankAccount{balance: 100}
	clienBank1.simulate(300)

	println(clienBank1.balance)

	leandro := Client{
		name: "Leandro",
	}

	leandro.walk()
	fmt.Printf("The value of the struct with name is this here %v", leandro.name)
}
