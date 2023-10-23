package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

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
	http.HandleFunc("/infinitescroll/userlist", userlist)

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

type LoadMore = struct {
	From int
	Size int
}

func infinitescroll(w http.ResponseWriter, r *http.Request) {
	t := template.Must(templateloader.GetBaseTemplate().ParseFiles("web/template/infinitescroll/infinitescroll.html"))

	t.ExecuteTemplate(w, "infinitescroll.html", struct {
		PageData
		Users []User
		LoadMore
	}{
		PageData: PageData{Title: "Infinite scroll"},
		Users:    getUserListPage(0, 10),
		LoadMore: LoadMore{
			From: 10,
			Size: 10,
		},
	})
}

func getUserListPage(from int, size int) []User {
	var res []User

	for i := from; i < from+size; i++ {
		res = append(res, User{
			Avatar: fmt.Sprintf("https://robohash.org/%d", i),
		})
	}

	return res
}

func userlist(w http.ResponseWriter, r *http.Request) {
	// simulate delay
	time.Sleep(1 * time.Second)

	var query = r.URL.Query()

	from, err := strconv.Atoi(query.Get("from"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<p>Invalid value for 'from' argument</p>"))
		return
	}

	size, err := strconv.Atoi(query.Get("size"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<p>Invalid value for 'size' argument</p>"))
		return
	}

	t := template.Must(templateloader.GetBaseTemplate().ParseFiles("web/template/infinitescroll/infinitescroll.html", "web/template/infinitescroll/userlist.html"))

	t.ExecuteTemplate(w, "userlist.html", struct {
		Users []User
		LoadMore
	}{
		Users: getUserListPage(from, size),
		LoadMore: LoadMore{
			From: from + size,
			Size: size,
		},
	})
}
