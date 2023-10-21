package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const PORT = 8000

func main() {
	fmt.Printf("Server started on http://localhost:%d/\n", PORT)

	// static files
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// hot reload
	http.HandleFunc("/_hotreloadws", hotreload)

	// // views
	http.HandleFunc("/", Homepage)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(
		"web/template/basepage.html",
	))

	t = template.Must(t.ParseFiles(
		"web/template/homepage.html",
	))

	t.Execute(w, struct{}{})
}
