package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	WorkLoad int
}

func main() {
	course1 := Course{"GO", 40}

	// Go has built-in template system, not need external libraries.
	template := template.New("CourseTemplate")

	template, error := template.Parse("Course: {{.Name}} - Carga Horária: {{.WorkLoad}}")

	if error != nil {
		panic(error)
	}

	error = template.Execute(os.Stdout, course1)

	if error != nil {
		panic(error)
	}
}
