package main

import (
	"errors"
	"fmt"
	"strings"
)

func containsString(collection []string, value string) bool {
	for _, v := range collection {
		if v == value {
			return true
		}
	}
	return false
}

func getSampleFieldFormatted(dataFields []string) ([]string, error) {
	var result []string

	mapObjects := map[string]string{
		"a": "tv",
		"b": "phone",
		"c": "computer",
		"d": "car",
		"e": "house",
		"f": "boat",
		"g": "airplane",
		"h": "motorcycle",
		"i": "bike",
		"j": "scooter",
	}

	for _, v := range dataFields {
		if _, ok := mapObjects[v]; ok {
			result = append(result, mapObjects[v])
		}
	}

	if containsString(dataFields, "a") || containsString(dataFields, "i") || containsString(dataFields, "j") {
		return nil, errors.New("column not found")
	}

	return result, nil
}

func searchBinary(slice []string) []string {
	sampleFields := []string{}

	fmt.Println("Trying Slice====>", slice)

	if len(slice) == 1 {
		return sampleFields
	}

	newSlices := splitSlice(slice)

	fmt.Println("newSlices", newSlices)

	for _, slice := range newSlices {

		fmt.Println("slice FOR", slice)

		if len(slice) == 1 {
			// pega o cache
			// if err != nil {
			// 		continue
		}

		data, err := getSampleFieldFormatted(slice)

		fmt.Println("data  ===>", data)

		if err == nil {
			sampleFields = append(sampleFields, data...)

			fmt.Println("sampleFields", sampleFields)

		} else {
			sampleFields = append(sampleFields, searchBinary(slice)...)

			fmt.Println("Search ====>", sampleFields)
		}
	}

	return sampleFields
}

func splitSlice(slice []string) [][]string {
	lenght := len(slice)
	middle := lenght / 2

	firstHalf := slice[:middle]
	secondHalf := slice[middle:]

	return [][]string{firstHalf, secondHalf}
}

func main() {
	expectedResult := []string{"phone", "computer", "car", "house", "boat", "airplane", "motorcycle"}
	result := searchBinary([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"})

	if strings.Join(expectedResult, ",") != strings.Join(result, ",") {
		fmt.Println(result)
		println("Error")
	} else {
		println("Success")
	}
}
