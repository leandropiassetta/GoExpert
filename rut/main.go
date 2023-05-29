package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func matchRutData(sample string) string {
	var allRuts string

	for i := 0; i < len(sample); i++ {
		initialPosition := i
		finalPosition := i + 12

		if finalPosition > len(sample) {
			break
		}

		rut := sample[initialPosition:finalPosition]

		rut, _, isRut := rutChile(rut)
		if !isRut {
			continue
		}

		allRuts += fmt.Sprintf("RUT ==>: %s\n", rut)
	}

	return allRuts
}

func rutChile(rut string) (string, error, bool) {
	// validate if rut have the correct format
	regexRut := regexp.MustCompile(`^[0-9]{2}\.[0-9]{3}\.[0-9]{3}-[0-9kK]$`)
	rutIsValid := regexRut.MatchString(rut)

	if !rutIsValid {
		fmt.Println("The rut is invalid:", rut)

		return "", fmt.Errorf("The rut is invalid: %s", rut), false
	}

	if !matchDigitVerificator(rut) {
		fmt.Println("The rut is invalid:", rut)

		return "", fmt.Errorf("The rut is invalid: %s", rut), false
	}

	fmt.Println("The rut is valid:", rut)

	return rut, nil, true
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
		digitVerificator = "k"
	}

	return digitVerificator
}

func matchDigitVerificator(rut string) bool {
	// remove the dots and the digit verificator from the rut to calculate the digit verificator again and compare with the digit verificator
	rutWithoutDots := strings.Split(strings.ReplaceAll(rut, ".", ""), "-")[0]

	// calculate the digit verificator
	digitVerificatorValid := calculateDigitVerificator(rutWithoutDots)

	// compare the digit verificator calculated with the digit verificator of the rut received
	digitVerificatorReceived := strings.Split(rut, "-")[1]

	if digitVerificatorValid != digitVerificatorReceived {
		return false
	}

	return true
}

func main() {
	// rutChile("20.216.023-9")
	// rutChile("64.630.536-5")
	// rutChile("25.566.862-5")
	// rutChile("42.733.846-0")
	// rutChile("86.962.200-1")
	// rutChile("37.311.245-3")
	// rutChile("35.203.080-5")
	// rutChile("10.812.817-8")
	// rutChile("83.452.556-9")
	// rutChile("30.510.663-1")
	// rutChile("16.505.024-k")
	// rutChile("16.505.24-k")
	// rutChile("00.000.24-k")
	// rutChile("99.999.99-9")
	// rutChile("11.111.11-1")
	// rutChile("20.216.aaa-9")
	sample := "20.216.023-9Lorem ipsum dolor sit amet, consectetur adipiscing elit rut: 20.216.023-9,lorem, ipsum dolor sit amet, consectetur adipiscing elit rut: 64.630.536-5, lorem ipsum dolor sit amet"
	fmt.Printf("Thats Ruts was founded in the sample: \n%s", matchRutData(sample))
}
