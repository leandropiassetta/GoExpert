package mock_test

import (
	"errors"
	"testing"

	"mock"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateTax_Success(t *testing.T) {
	tax, err := mock.CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax, "Tax should be 10.0")
}

func Test_CalculateTax_Error(t *testing.T) {
	tax, err := mock.CalculateTax(0)

	assert.Error(t, err)
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), " greater than 0")
}

func Test_CalculateTaxAndSave_Success(t *testing.T) {
	// now i can simulate the SaveTax method call and the argument passed to the SaveTax method call (i can simulate the SaveTax method call and the argument passed to the SaveTax method call because i have a mock implementation of the Repository interface)
	repositoryMock := &mock.TaxRepositoryMock{}

	// mock.Mock.On -> is a method that returns a struct that has a field called "Arguments" that is a slice of interface

	repositoryMock.On("SaveTax", 10.0).Return(nil)

	err := mock.CalculateTaxAndSave(1000.0, repositoryMock)

	assert.Nil(t, err)
}

func Test_CalculateTaxAndSave_Error(t *testing.T) {
	repositoryMock := &mock.TaxRepositoryMock{}

	repositoryMock.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := mock.CalculateTaxAndSave(0, repositoryMock)

	assert.Error(t, err)
}

// test with AssertExpectations(t)
func Test_CalculateTaxAndSave_Assertions(t *testing.T) {
	repositoryMock := &mock.TaxRepositoryMock{}

	repositoryMock.On("SaveTax", 10.0).Return(nil).Twice()
	repositoryMock.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	// repositoryMock.On("SaveTax", mock.Anything).Return(errors.New("error saving tax"))

	err := mock.CalculateTaxAndSave(1000.0, repositoryMock)
	assert.Nil(t, err)

	err = mock.CalculateTaxAndSave(1000.0, repositoryMock)
	assert.Nil(t, err)

	err = mock.CalculateTaxAndSave(0, repositoryMock)

	assert.Error(t, err, "error saving tax")

	// mock.Mock.AssertExpectations -> is a method that returns a bool (true if all the method calls and arguments passed to the method calls were simulated, false otherwise) (it is used to test the code that uses the Repository interface)
	repositoryMock.AssertExpectations(t)

	// mock.Mock.AssertNumberOfCalls -> is a method that returns a bool (true if the method call was simulated the number of times passed to the method, false otherwise) (it is used to test the code that uses the Repository interface)
	repositoryMock.AssertNumberOfCalls(t, "SaveTax", 3)
}
