package mathematic

// Constraint that i created
type Number interface {
	~int | ~float64
}

// For export one function of my package i have that write with the first letter of my function in the uppercase, for utilize this function only in this package i write the function with the first letter in the lowercase.
func Sum[T Number](numA, numB T) T {
	return numA + numB
}
