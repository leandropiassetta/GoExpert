package user

// Entity is where we define our rules of our business logic and our data structure for our application.

import (
	"github.com/leandropiassetta/goexpert/9-apis/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// VO - Value Object
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // the minus sign means that this field will not be marshalled and unmarshalled from the json file (it will not be exposed to the outside world) but it will be used internally in our application (for example to compare the password with the hash password) and it will be stored in the database (because we need to store the hash password in the database)
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}, nil
}

// ValidatePassword is a method that will be used to compare the password with the hash password
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
