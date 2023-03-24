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

	// I can execute this lines of code in one command, i can:
	// template := template.New("CourseTemplate")
	// template, _ := template.Parse("Course: {{.Name}} - Carga Horária: {{.WorkLoad}}")

	// and if there is already an error it already returns this error for me.

	// When we use a recurse called of Template.Must

	template := template.Must(template.New("Curso Template").Parse("Course: {{.Name}} - Carga Horária: {{.WorkLoad}}"))
	error := template.Execute(os.Stdout, course1)

	if error != nil {
		panic(error)
	}
}
