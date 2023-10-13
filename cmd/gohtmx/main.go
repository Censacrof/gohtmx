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
	fmt.Println("Server started")

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
