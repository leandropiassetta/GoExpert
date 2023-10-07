package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateTax_Success(t *testing.T) {
	tax, err := CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax, "Tax should be 10.0")
}

func Test_CalculateTax_Error(t *testing.T) {
	tax, err := CalculateTax(0)

	assert.Error(t, err)
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), " greater than 0")
}
