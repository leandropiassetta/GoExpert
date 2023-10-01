package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/leandropiassetta/goexpert/9-apis/configs"
	"github.com/leandropiassetta/goexpert/9-apis/internal/entity"
	"github.com/leandropiassetta/goexpert/9-apis/internal/infra/database"
	"github.com/leandropiassetta/goexpert/9-apis/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// its necessary to import the docs package to generate the swagger docs
	_ "github.com/leandropiassetta/goexpert/9-apis/docs"
)

// @title Swagger Go Expert API Example
// @version 1.0
// @description Product API Example with a JWT authentication
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	userHandler := handlers.NewUserHandler(userDB)

	// create a new router and register the handler functions for each route defined
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// router.Use(LogRequest)

	// recoverer middleware - recover from panics without crashing the server, the web server will continue running
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", configs.TokenAuth))
	router.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExperesIn))

	// POST /products - create a new product and return the product in the response body as JSON (StatusCreated) or return a status code 400 (BadRequest) if the request body is invalid or a status code 500 (InternalServerError) if there was an error while saving the product into the database

	// Product
	router.Route("/products", func(router chi.Router) {
		// jwt middleware - protected routes - only authenticated users can access
		router.Use(jwtauth.Verifier(configs.TokenAuth))
		router.Use(jwtauth.Authenticator)
		router.Post("/", productHandler.CreateProduct)
		router.Get("/", productHandler.GetProducts)
		router.Get("/{id}", productHandler.GetProduct)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.DeleteProduct)
	})

	// User
	router.Post("/users", userHandler.Create)
	router.Post("/users/generate_token", userHandler.GetJWT)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	println("Server running on port 8000")
	// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux
	http.ListenAndServe(":8000", router)
}

// client -> request -> http handler -> middleware -> http handler -> response -> client

// the middleware receive a paramater that will be called next. This parameter is a function that will be called after the middleware logic is executed, will continue your execution for the next middleware or the handler function
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Context().Value("user")

		log.Printf("Request: %s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
