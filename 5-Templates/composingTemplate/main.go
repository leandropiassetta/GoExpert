package main

import (

	// have two package for template html and text, what the difference between the two?

	// html -> he know that parse in html and knowing this will shield itself from attacks, its better for show in the browser
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name     string
	WorkLoad int
}

type Courses []Course

func main() {
	// to make Parse for many files

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")

	// i can make a map of functions here inside to be parsed
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(template.ParseFiles(templates...))

	error := t.Execute(os.Stdout, Courses{
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
