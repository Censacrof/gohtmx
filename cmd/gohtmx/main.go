package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Censacrof/gohtmx/cmd/gohtmx/view"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/homepage.html")

	t.Execute(w, struct{}{})
}

func main() {
	fmt.Println("Server started on http://localhost:8080/")

	http.HandleFunc("/", view.Homepage)

	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
