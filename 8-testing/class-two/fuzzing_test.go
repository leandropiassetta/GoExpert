package fuzzing

import "testing"

// fuzzing -> testing the code with random data (it's a good way to find bugs in the code)

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 0, 1, 2, 2.5, 500.0, 1000.0, 1501.0, 2000.0, 2001.0, 5000.0, 10000.0, 20000.0, 50000.0, 100000.0, 200000.0, 500000.0, 1000000.0}

	for _, amount := range seed {
		// f.Fuzz(amount) -> it's a good way to find bugs in the code
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}

		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})
}

// how to run fuzzing: go test -fuzz=FuzzCalculateTax -count=1000
// -count=1000 -> how many times to run the fuzzing

// go test -fuzz=. -run=^# -> run all fuzzing tests (it's a good way to find bugs in the code)
// go test -fuzz=FuzzCalculateTax -count=1000 -cover
// go test -fuzz=FuzzCalculateTax -count=1000 -cover -coverprofile=coverage.out
// go tool cover -html=coverage.out
// go test -fuzz=FuzzCalculateTax -count=1000 -cover -coverprofile=coverage.out -covermode=count
