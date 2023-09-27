package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/leandropiassetta/goexpert/9-apis/configs"
	"github.com/leandropiassetta/goexpert/9-apis/internal/entity"
	"github.com/leandropiassetta/goexpert/9-apis/internal/infra/database"
	"github.com/leandropiassetta/goexpert/9-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// config := configs.NewConfig()
	// println(config.GetDbDriver())

	configs, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	// println(config.DBDriver)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	// Product
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	// User

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExperesIn)

	// create a new router and register the handler functions for each route defined
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	println("Server running on port 8000")

	// POST /products - create a new product and return the product in the response body as JSON (StatusCreated) or return a status code 400 (BadRequest) if the request body is invalid or a status code 500 (InternalServerError) if there was an error while saving the product into the database
	// Product
	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products", productHandler.GetProducts)
	router.Get("/products/{id}", productHandler.GetProduct)
	router.Put("/products/{id}", productHandler.UpdateProduct)
	router.Delete("/products/{id}", productHandler.DeleteProduct)

	// User
	router.Post("/users", userHandler.Create)
	router.Post("/users/generate_token", userHandler.GetJWT)

	// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux
	http.ListenAndServe(":8000", router)
}
