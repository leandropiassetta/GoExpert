package entity_test

import (
	"testing"

	"github.com/leandropiassetta/goexpert/9-apis/internal/entity"
	"github.com/stretchr/testify/assert"
)

func Test_NewProduct(t *testing.T) {
	p, err := entity.NewProduct("Product 1", 100)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 100, p.Price)
	assert.NotEmpty(t, p.CreatedAt)
}

func Test_ProductWhenNameIsRequired(t *testing.T) {
	p, err := entity.NewProduct("", 100)

	assert.Nil(t, p)
	assert.Equal(t, entity.ErrorNameIsRequired, err)
}

func Test_ProductWhenPriceIsRequired(t *testing.T) {
	p, err := entity.NewProduct("Product 1", 0)

	assert.Nil(t, p)
	assert.Equal(t, entity.ErrorPriceIsRequired, err)
}

func Test_ProductWhenPriceIsInvalid(t *testing.T) {
	p, err := entity.NewProduct("Product 1", -1)

	assert.Nil(t, p)
	assert.Equal(t, entity.ErrorInvalidPrice, err)
}

func Test_ProductValidate(t *testing.T) {
	p, err := entity.NewProduct("Product 1", 10)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
