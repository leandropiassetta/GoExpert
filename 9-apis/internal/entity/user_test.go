package entity_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/leandropiassetta/goexpert/9-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Bcrypt interface {
	CompareHashAndPassword(hashedPassword, password []byte) error
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
}

type MockBcrypt struct {
	mock.Mock
}

func (m *MockBcrypt) CompareHashAndPassword(hashedPassword, password []byte) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}

type UserMock struct {
	mock.Mock
}

func (user *UserMock) NewUser(name, email, password string) (entity.User, error) {
	mockBcrypt := &MockBcrypt{}

	if len(password) > 72 {
		mockBcrypt.On("GenerateFromPassword", []byte(password), 10).Return([]byte{}, errors.New("password is too long"))
	} else {
		expectedHash := []byte(password)
		mockBcrypt.On("GenerateFromPassword", []byte(password), 10).Return(expectedHash, nil)
	}

	args := user.Called(name, email, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (user *UserMock) ValidatePassword(password string) bool {
	args := user.Called(password)
	return args.Bool(0)
}

func Test_NewUser(t *testing.T) {
	userMock := &UserMock{}
	id := uuid.UUID{}

	expectedUser := entity.User{
		ID:       id,
		Name:     "Leandro",
		Email:    "leandro@gmail.com",
		Password: "123456",
	}

	userMock.On("NewUser", "Leandro", "leandro@gmail.com", "123456").Return(expectedUser, nil)

	newUser, err := userMock.NewUser("Leandro", "leandro@gmail.com", "123456")

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, newUser)
}
