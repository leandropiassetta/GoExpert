package tax

func CalculateTax(amount float64) float64 {
	tax := 5.0

	if amount >= 1000 {
		tax = 10.0
		return tax
	}

	return tax
}
