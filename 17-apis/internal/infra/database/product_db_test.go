package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/leandropiassetta/goexpert/09-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_CreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)

	productDB := NewProduct(db)

	err = productDB.CreateProduct(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func Test_FindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product: %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")

	assert.NoError(t, err)
	assert.Equal(t, 10, len(products))
	// assert.Subset(t, []string{"Product: 1", "Product: 2", "Product: 3", "Product: 4", "Product: 5", "Product: 6", "Product: 7", "Product: 8", "Product: 9", "Product: 10"}, []string{products[0].Name, products[1].Name, products[2].Name, products[3].Name, products[4].Name, products[5].Name, products[6].Name, products[7].Name, products[8].Name, products[9].Name})
	assert.Equal(t, "Product: 1", products[0].Name)
	assert.Equal(t, "Product: 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product: 11", products[0].Name)
	assert.Equal(t, "Product: 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(products))
	assert.Equal(t, "Product: 21", products[0].Name)
	assert.Equal(t, "Product: 24", products[3].Name)
}

func Test_FindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindProductByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func Test_UpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)

	product.Name = "Product 2"
	product.Price = 20.00
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindProductByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func Test_DeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindProductByID(product.ID.String())
	assert.Error(t, err)
}
