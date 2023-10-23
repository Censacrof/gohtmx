package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const PORT = 8000

func main() {
	fmt.Printf("Server started on http://localhost:%d/\n", PORT)

	// static files
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// hot reload
	http.HandleFunc("/_hotreloadws", hotreload)

	// views
	http.HandleFunc("/", homepage)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/template/homepage.html"))

	t.Execute(w, struct{}{})
}
