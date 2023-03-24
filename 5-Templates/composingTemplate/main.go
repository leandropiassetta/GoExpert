package main

import (
	"net/http"
	"text/template"
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

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		template := template.Must(template.New("content.html").ParseFiles(templates...))

		// get the datas of my execute and throw in my response writer.
		error := template.Execute(writer, Courses{
			{"GO", 140},
			{"JavaScript", 80},
			{"Python", 30},
			{"React", 50},
			{"Node", 100},
		})

		if error != nil {
			panic(error)
		}
	})

	http.ListenAndServe(":8282", nil)
}
