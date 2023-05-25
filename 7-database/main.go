package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		// i want that to assign only once the id

		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	fmt.Println()
}
