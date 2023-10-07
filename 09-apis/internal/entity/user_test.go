package entity_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/leandropiassetta/goexpert/09-apis/internal/entity"
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

func (u *UserMock) NewUser(name, email, password string) (entity.User, error) {
	mockBcrypt := &MockBcrypt{}

	if len(password) > 72 {
		mockBcrypt.On("GenerateFromPassword", []byte(password), 10).Return([]byte{}, errors.New("password is too long"))
	} else {
		expectedHash := []byte(password)
		mockBcrypt.On("GenerateFromPassword", []byte(password), 10).Return(expectedHash, nil)
	}

	args := u.Called(name, email, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserMock) ValidatePassword(password string) bool {
	args := u.Called(password)
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

	t.Run("should return an error if password is too long", func(t *testing.T) {
		userMock.On("NewUser", "Leandro", "leandro@gmail.com", "senha_muito_longa_que_vai_ultrapassar_o_limite_de_72_caracteres_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").Return(entity.User{}, errors.New("password is too long"))

		_, err := userMock.NewUser("Leandro", "leandro@gmail.com", "senha_muito_longa_que_vai_ultrapassar_o_limite_de_72_caracteres_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

		assert.Error(t, err)
	})

	t.Run("should return a user", func(t *testing.T) {
		userMock.On("NewUser", "Leandro", "leandro@gmail.com", "123456").Return(expectedUser, nil)
	})

	user, err := userMock.NewUser("Leandro", "leandro@gmail.com", "123456")

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func Test_User_Other_Way(t *testing.T) {
	user, err := entity.NewUser("Leandro", "l@gmail.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Leandro", user.Name)
	assert.Equal(t, "l@gmail.com", user.Email)
}

func Test_User_ValidatePassword(t *testing.T) {
	user, err := entity.NewUser("Leandro", "l@gmail.com", "123456")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
