package mock

import "errors"

// interface -> is a contract that a type can implement (by defining the methods in the interface) i garantees that the type will have those methods defined in the interface (the contract)

type Repository interface {
	SaveTax(amount float64) error
}

// mock -> is a fake implementation of an interface, it is used to test the code that uses the interface, it is used to test the code that uses the interface (the contract)

// i want to test the CalculateTaxAndSave function, but it uses the Repository interface, so i need to create a mock implementation of the Repository interface
// and isolate the CalculateTaxAndSave function from the real implementation of the Repository interface (the real implementation of the Repository interface will be used in production) and use the mock implementation of the Repository interface in the tests

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax_2(amount)

	return repository.SaveTax(tax)
}

func CalculateTax_2(amount float64) float64 {
	tax := 5.0

	if amount <= 0 {
		tax = 0.0
		return tax
	}

	if amount >= 1000 && amount < 20000 {
		tax = 10.0
		return tax

	}

	if amount >= 20000 {
		tax = 20.0
		return tax

	}

	return tax
}

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
