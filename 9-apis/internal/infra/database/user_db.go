package database

import (
	"github.com/leandropiassetta/goexpert/9-apis/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) CreateUser(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	// this is the same as SELECT * FROM users WHERE email = ?
	// &user is the pointer to the user variable that we want to fill with the data from the database
	err := u.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
