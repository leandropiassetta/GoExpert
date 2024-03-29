package entity

import (
	"errors"
	"time"

	"github.com/leandropiassetta/goexpert/09-apis/pkg/entity"
)

var (
	ErrorIDIsRequired    = errors.New("id is required")
	ErrorInavalidID      = errors.New("invalid id")
	ErrorNameIsRequired  = errors.New("name is required")
	ErrorPriceIsRequired = errors.New("price is required")
	ErrorInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorInavalidID
	}

	if p.Name == "" {
		return ErrorNameIsRequired
	}

	if p.Price == 0.00 {
		return ErrorPriceIsRequired
	}

	if p.Price < 0.00 {
		return ErrorInvalidPrice
	}

	return nil
}
