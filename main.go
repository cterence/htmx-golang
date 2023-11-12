package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Film struct {
	Title string
	Year  int
}

func main() {
	fmt.Println("HTMX test")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("src/index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Casablanca", Year: 1942},
				{Title: "Cool Hand Luke", Year: 1967},
				{Title: "Bullitt", Year: 1968},
				{Title: "Interstellar", Year: 2014},
				{Title: "Mustang", Year: 2015},
				{Title: "Demolition", Year: 2015},
				{Title: "Avengers: Endgame", Year: 2019},
			},
		}

		tmpl.Execute(w, films)
	})

	http.HandleFunc("/film/add", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		log.Print("Add film")
		title := r.PostFormValue("title")
		year := r.PostFormValue("year")
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			log.Fatal(err)
		}
		tmpl := template.Must(template.ParseFiles("src/index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Year: yearInt})
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
