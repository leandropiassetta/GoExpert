package tax

import "testing"

// go test -v -> verbose mode (shows the name of the tests)

func Test_CalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected: %f, but got: %f", expected, result)
	}
}
