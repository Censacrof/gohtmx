package view

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Censacrof/gohtmx/cmd/gohtmx/components"
)

func getUserListEntriesPage(from int, size int) template.HTML {
	var entries template.HTML

	for i := 0; i < size; i++ {
		avatar := template.URL(fmt.Sprintf("https://robohash.org/%d", i))
		entries += components.UserListEntry{
			Avatar: avatar,
		}.HTML()
	}
	return entries
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	intial_entries := getUserListEntriesPage(0, 10)

	var body = components.BasePage{
		Title: "Homepage",
		Body: components.Wrapper{
			Content: components.UserList{
				Entries: intial_entries,
			}.HTML(),
		}.HTML(),
	}.HTML()

	w.Write([]byte(body))
}
