package testify

import "errors"

func CalculateTax(amount float64) (float64, error) {
	tax := 5.0

	if amount <= 0 {
		tax = 0.0
		return tax, errors.New("amount should be greater than 0")
	}

	if amount == 0 {
		tax = 0.0
		return tax, nil
	}

	if amount >= 1000 && amount < 20000 {
		tax = 10.0
		return tax, nil

	}

	if amount >= 20000 {
		tax = 20.0
		return tax, nil

	}

	return tax, nil
}
