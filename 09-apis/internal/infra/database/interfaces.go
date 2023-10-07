package database

import "github.com/leandropiassetta/goexpert/09-apis/internal/entity"

// this is the interface that we will use to interact with the database, my application less coupled trying to use the interface instead of the implementation directly in the code

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	CreateProduct(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindProductByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
