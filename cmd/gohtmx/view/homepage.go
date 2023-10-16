package view

import (
	"net/http"

	"github.com/Censacrof/gohtmx/cmd/gohtmx/components"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	var body = components.BasePage{
		Title: "Homepage",
		Body: components.Wrapper{
			Content: components.UserList{
				Entries: components.UserListEntry{
					Avatar: "https://robohash.org/0",
				}.HTML() + components.UserListEntry{
					Avatar: "https://robohash.org/1",
				}.HTML(),
			}.HTML(),
		}.HTML(),
	}.HTML()

	w.Write([]byte(body))
}
