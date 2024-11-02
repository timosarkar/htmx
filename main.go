package main

import (
"fmt"
	"os"
	_ "github.com/glebarez/go-sqlite"
	"database/sql"
	"html/template"
	"net/http"
)

func createDB() error {
	db, err := sql.Open("sqlite", "index.db")
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}

func main() {
if err := createDB(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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


//	http.ListenAndServe(":8080", nil)
}

