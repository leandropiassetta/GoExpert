package entity

// Entity is where we define our rules of our business logic and our data structure for our application.

import (
	"github.com/x/crypto"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // the minus sign means that this field will not be marshalled and unmarshalled from the json file (it will not be exposed to the outside world) but it will be used internally in our application (for example to compare the password with the hash password) and it will be stored in the database (because we need to store the hash password in the database)
}

func NewUser(id, name, email, password string) (*User, error) {
	id = bcrypt.GenerateID()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
}
