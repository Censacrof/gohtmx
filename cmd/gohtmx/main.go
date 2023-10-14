package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/index.html")

	page := &Page{
		Title: "Homepage",
	}

	t.Execute(w, page)
}

func main() {
	fmt.Println("Server started on http://localhost:8080/")

	http.HandleFunc("/", handler)

	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
