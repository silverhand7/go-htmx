package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello world")
	h1 := func(w http.ResponseWriter, r *http.Request) {
		films := map[string][]Film{
			"Films": {
				{Title: "Breaking Bad", Director: "Vince Gilligan"},
				{Title: "Game of Thrones", Director: "George RR Martin"},
				{Title: "Peaky Blinders", Director: "Steven Kinght"},
			},
		}
		pageTemplate := template.Must(template.ParseFiles("index.html"))
		pageTemplate.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf(`<li class="list-group-item bg-primary text-white">%s - %s</li>`, title, director)
		pageTemplate, _ := template.New("t").Parse(htmlStr)
		pageTemplate.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
