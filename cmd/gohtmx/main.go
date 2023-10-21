package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Censacrof/gohtmx/cmd/gohtmx/templateloader"
	"github.com/a-h/templ"
)

const PORT = 8000

func main() {
	fmt.Printf("Server started on http://localhost:%d/\n", PORT)

	component := hello("John")

	http.Handle("/", templ.Handler(component))

	// static files
	// fs := http.FileServer(http.Dir("./web/static"))
	// http.Handle("/static/", http.StripPrefix("/static", fs))

	// // hot reload
	// http.HandleFunc("/_hotreloadws", hotreload)

	// // // views
	// http.HandleFunc("/", Homepage)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	t := templateloader.GetTemplate()

	t.ExecuteTemplate(w, "homepage.view.html", struct{}{})
}
