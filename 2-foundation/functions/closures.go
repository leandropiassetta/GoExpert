package main

import (
	"fmt"
	"regexp"
)

// uma função retorna o valor de outra função

func main() {
	total := func() int {
		return sum(1, 3, 4, 5, 6) * 2
	}()

	fmt.Println(total)

	fmt.Println(removeVowels("qualquer cOisA"))
	fmt.Println(vowelPickup("qualquer cOisA"))
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func vowelPickup(phrase string) []string {
	regexStr := regexp.MustCompile(`[aeiouAEIOU]`)
	fmt.Println(regexStr.MatchString(phrase))

	allVowels := regexStr.FindAllString(phrase, -1)
	fmt.Println("Number of vowels: ", len(allVowels))
	return allVowels
}

func removeVowels(phrase string) string {
	matchVowel := regexp.MustCompile(`[aeiouAEIOU]`)
	noVowel := matchVowel.ReplaceAllString(phrase, "")

	return noVowel
}
