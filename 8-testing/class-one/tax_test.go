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

// batch -> many tests in one file
func Test_CalculateTax_Batch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{amount: 500, expect: 5},
		{amount: 1000, expect: 10},
		{amount: 1500, expect: 10},
		// {amount: 2000, expect: 12},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expect {
			t.Errorf("Expected: %f, but got: %f", item.expect, result)
		}
	}
}
