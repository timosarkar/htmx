package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("index.templ"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Determine which partial to load based on route
		var section string
		switch r.URL.Path {
		case "/features":
			section = "features"
		case "/contact":
			section = "contact"
		default:
			section = "home"
		}

		tmpl.Execute(w, map[string]string{"Section": section})
	})


	http.ListenAndServe(":8080", nil)
}

