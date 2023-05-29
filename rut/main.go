package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// matchRutData find all ruts in the sample and check if the rut is valid
func matchRutData(sample string) string {
	var allRuts string
	regexRut := regexp.MustCompile(`\d{2}[\s\.]?\d{3}[\s\.]?\d{3}[\s\-]?[0-9kK]`)
	matches := regexRut.FindAllString(sample, -1)

	fmt.Printf("ruts to be checked: %v\n", matches)

	for _, rut := range matches {

		isRut := matchDigitVerificator(rut)
		if isRut {
			allRuts += fmt.Sprintf("rut found: %s\n", rut)
		}
	}

	return allRuts
}

func matchDigitVerificator(rut string) bool {
	rutWithoutSymbols := cleanRut(rut)

	digitVerificatorRutReceived := rutWithoutSymbols[len(rutWithoutSymbols)-1:]
	rutReceivedWithoutDigitVerificator := rutWithoutSymbols[:len(rutWithoutSymbols)-1]

	digitVerificatorValid := calculateDigitVerificator(rutReceivedWithoutDigitVerificator)

	fmt.Println("digit verificator valid: ", digitVerificatorValid)

	// compare the digit verificator valid with the digit verificator of the rut received
	if strings.ToUpper(digitVerificatorValid) != strings.ToUpper(digitVerificatorRutReceived) {
		fmt.Printf("The rut is invalid: %s ", rut)
		return false
	}

	return true
}

// cleanRut remove the symbols of the rut received for calculate the digit verificator valid and compare with the digit verificator of the rut received
func cleanRut(rut string) string {
	rutWithoutSymbols := strings.ReplaceAll(rut, ".", "")
	rutWithoutSymbols = strings.ReplaceAll(rutWithoutSymbols, "-", "")
	rutWithoutSymbols = strings.ReplaceAll(rutWithoutSymbols, " ", "")

	return rutWithoutSymbols
}

func calculateDigitVerificator(rutWithoutDots string) string {
	multiplier := 2
	sum := 0

	for i := len(rutWithoutDots) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(rutWithoutDots[i]))
		sum += digit * multiplier

		// for the correct calculation of the digit verificator the multiplier must be between 2 and 7, if pass 7 the multiplier must be 2 again
		multiplier++
		if multiplier > 7 {
			multiplier = 2
		}
	}

	// apply the module 11
	moduleEleven := 11 - (sum % 11)
	digitVerificator := strconv.Itoa(moduleEleven)

	if digitVerificator == "11" {
		digitVerificator = "0"
	} else if digitVerificator == "10" {
		digitVerificator = "K"
	}

	return digitVerificator
}

func main() {
	sample := "Lorem ipsum dolor sit amet, consectetur adipiscing elit rut:53 723 007 K,lorem, ipsum 	 dolor sit amet, consectetur adipiscing eli, lorem ipsum dolor sit amet"
	fmt.Printf("these ruts were found in the samples: \n%s", matchRutData(sample))
}
