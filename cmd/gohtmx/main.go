package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Censacrof/gohtmx/cmd/gohtmx/templateloader"
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
	http.HandleFunc("/infinitescroll/", infinitescroll)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

type PageData = struct {
	Title string
}

func homepage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(templateloader.GetBaseTemplate().ParseFiles("web/template/homepage.html"))

	t.ExecuteTemplate(w, "homepage.html", PageData{
		Title: "Hompage",
	})
}

type User = struct {
	Avatar string
}

func infinitescroll(w http.ResponseWriter, r *http.Request) {
	t := template.Must(templateloader.GetBaseTemplate().ParseFiles("web/template/infinitescroll/infinitescroll.html"))

	t.ExecuteTemplate(w, "infinitescroll.html", struct {
		PageData
		Users []User
	}{
		PageData: PageData{Title: "Infinite scroll"},
		Users: []User{
			{
				Avatar: "https://robohash.org/0",
			},
			{
				Avatar: "https://robohash.org/1",
			},
			{
				Avatar: "https://robohash.org/2",
			},
		},
	})
}
