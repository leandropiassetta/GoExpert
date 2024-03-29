package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/leandropiassetta/goexpert/09-apis/internal/dto"
	"github.com/leandropiassetta/goexpert/09-apis/internal/entity"
	"github.com/leandropiassetta/goexpert/09-apis/internal/infra/database"
	entityPkg "github.com/leandropiassetta/goexpert/09-apis/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct godoc
// @Summary Create a product
// @Description Create a product with the input payload
// @Tags products
// @Accept  json
// @Produce  json
// @Param request body dto.CreateProductInput true "product request"
// @Success 201
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
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

// GetProduct godoc
// @Summary Get a product
// @Description Get a product by id
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "product id" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 404
// @Failure 500 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the product from the database
	product, err := h.ProductDB.FindProductByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// return the product in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "product id" Format(uuid)
// @Param request body dto.CreateProductInput true "product request"
// @Success 200
// @Failure 404
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	var product entity.Product
	// decode the request body into product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindProductByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// update the product in the database with the new values from the request body (product)
	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type Response struct {
		Message string      `json:"message"`
		Product interface{} `json:"product"`
	}

	// return the product in the response

	response := Response{
		Message: "Product updated successfully",
		Product: product,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by id
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "product id" Format(uuid)
// @Success 200
// @Failure 404
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	product, err := h.ProductDB.FindProductByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// delete the product from the database
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("this product with id %s was deleted", product.ID))
}

// ListProducts godoc
// @Summary List products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Param page query string false "page number"
// @Param limit query string false "limit number"
// @Success 200 {array} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	// get all products from the database
	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return the products in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
