package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Censacrof/gohtmx/cmd/gohtmx/view"
)

const PORT = 8000

func main() {
	fmt.Printf("Server started on http://localhost:%d/\n", PORT)

	http.HandleFunc("/_hotreloadws", hotreload)

	http.HandleFunc("/", view.Homepage)

	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
