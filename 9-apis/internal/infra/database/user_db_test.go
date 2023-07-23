package database

import (
	"testing"

	"github.com/leandropiassetta/goexpert/9-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_CreateUser(t *testing.T) {
	// here we are creating a stub database connection in memory to run our tests without the need to connect to a real database like MySQL or Postgres for example (this is a good practice)
	// the stub database connection is created in memory and it is destroyed when the test finishes
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// this is the same as CREATE TABLE users (id integer primary key autoincrement, email text, password text, created_at datetime, updated_at datetime);
	// we are creating the table users in the database
	// Automigrate() -> is a gorm function that creates the table if it does not exist
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.CreateUser(user)
	assert.Nil(t, err)

	var userFound *entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.NotNil(t, userFound.Password)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.ID, userFound.ID)
}

func Test_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.CreateUser(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Password, userFound.Password)
}
