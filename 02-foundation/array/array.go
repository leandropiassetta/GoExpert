package main

import (
	"fmt"
	"strings"
)

const a = "hello world"

type ID int

var (
	b bool   = true
	c int    = 10
	d string = "Leandro"
	f ID     = 1
)

type DataPii struct {
	Content string
	Int     int
}

type SampleData struct {
	DataName    string
	Sample      string
	SampleCount int
}

func parseData(content string) []SampleData {
	content = strings.ReplaceAll(content, ", ", ",")
	contentParts := strings.Split(content, ", ")
	var sampleDataList []SampleData
	var currentSampleData SampleData

	for _, part := range contentParts {
		keyValue := strings.Split(part, " : ")
		if len(keyValue) != 2 {
			continue
		}

		key, value := strings.TrimSpace(keyValue[0]), strings.TrimSpace(keyValue[1])

		if strings.ContainsAny(key, " :") {
			if currentSampleData.DataName != "" {
				sampleDataList = append(sampleDataList, currentSampleData)
			}
			currentSampleData = SampleData{
				DataName:    key,
				Sample:      value,
				SampleCount: 1,
			}
		} else {
			currentSampleData.Sample += ", " + value
			currentSampleData.SampleCount++
		}
	}
	sampleDataList = append(sampleDataList, currentSampleData)

	return sampleDataList
}

func main() {
	fmt.Printf("O tipo de E é %T", f)

	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	// ultimo elemento dinamico

	fmt.Println(meuArray[len(meuArray)-1])

	// percorrendo Array

	for i, v := range meuArray {
		fmt.Printf("o valor do indice %d é %d\n", i, v)
	}

	data := DataPii{
		Content: "id : 1, 2, infotype : PERSON_NAME, EMAIL",
		Int:     2,
	}

	sampleDataList := parseData(data.Content)

	fmt.Println(sampleDataList)

	for _, sampleData := range sampleDataList {
		fmt.Printf("DataName: %s\n", sampleData.DataName)
		fmt.Printf("Sample: %s\n", sampleData.Sample)
		fmt.Printf("SampleCount: %d\n\n", sampleData.SampleCount)
	}
}
