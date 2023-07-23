package database

import "github.com/leandropiassetta/goexpert/9-apis/internal/entity"

// this is the interface that we will use to interact with the database, my application less coupled trying to use the interface instead of the implementation directly in the code

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
