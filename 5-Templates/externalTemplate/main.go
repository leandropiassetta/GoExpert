package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	WorkLoad int
}

type Courses []Course

func main() {
	template := template.Must(template.New("template.html").ParseFiles("template.html"))
	error := template.Execute(os.Stdout, Courses{
		{"GO", 140},
		{"JavaScript", 80},
		{"Python", 30},
		{"React", 50},
		{"Node", 100},
	})

	if error != nil {
		panic(error)
	}
}
