package view

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/Censacrof/gohtmx/cmd/gohtmx/components/userlist"
)

func getUserListEntriesPage(from int, size int) template.HTML {
	var entries template.HTML

	for i := 0; i < size; i++ {
		avatar := template.URL(fmt.Sprintf("https://robohash.org/%d", from+i))
		entries += userlist.UserListEntry{
			Avatar: avatar,
		}.HTML()
	}

	entries += userlist.LoadMore{
		From: from + size,
		Size: size,
	}.HTML()

	return entries
}

func UserListEntries(w http.ResponseWriter, r *http.Request) {
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

	w.Write([]byte(getUserListEntriesPage(from, size)))
}
