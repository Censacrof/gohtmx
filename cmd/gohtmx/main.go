package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

const PORT = 8000

func main() {
	fmt.Printf("Server started on http://localhost:%d/\n", PORT)

	// static files
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// hot reload
	http.HandleFunc("/_hotreloadws", hotreload)

	// // // views
	http.Handle("/", templ.Handler(HomePage()))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
