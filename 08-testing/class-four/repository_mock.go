package mock

import (
	"github.com/stretchr/testify/mock"
)

// mock.Mock -> is a struct that implements the Repository interface (the contract) and has a field called "Mock" that is a struct that has a field called "Calls" that is a slice of struct that has a field called "Arguments" that is a slice of interface
// mock.Mock simulates us method calls and arguments passed to the method calls (it is used to test the code that uses the Repository interface)
type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	// mock.Mock.Called -> is a method that returns a struct that has a field called "Arguments" that is a slice of interface
	args := m.Called(tax)
	return args.Error(0)
}
