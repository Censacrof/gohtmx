package view

import (
	"net/http"

	"github.com/Censacrof/gohtmx/cmd/gohtmx/components"
)

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
