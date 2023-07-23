package main

import (
	"encoding/json"
	"net/http"

	"github.com/leandropiassetta/goexpert/9-apis/configs"
	"github.com/leandropiassetta/goexpert/9-apis/internal/dto"
	"github.com/leandropiassetta/goexpert/9-apis/internal/entity"
	"github.com/leandropiassetta/goexpert/9-apis/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// config := configs.NewConfig()
	// println(config.GetDbDriver())

	config, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	println(config.DBDriver)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	println(db)

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	http.ListenAndServe(":8000", nil)
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		ProductDB: database.NewProduct(db),
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	// decode the request body into product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create a new product entity
	newProduct, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// save the product into the database
	err = h.ProductDB.CreateProduct(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return the product in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
