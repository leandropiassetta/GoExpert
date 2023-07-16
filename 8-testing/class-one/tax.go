package tax

import "time"

func CalculateTax(amount float64) float64 {
	tax := 5.0

	if amount == 0 {
		tax = 0.0
		return tax
	}

	if amount >= 1000 {
		tax = 10.0
		return tax
	}

	return tax
}

func CalculateTax_2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	tax := 5.0

	if amount == 0 {
		tax = 0.0
		return tax
	}

	if amount >= 1000 {
		tax = 10.0
		return tax
	}

	return tax
}
