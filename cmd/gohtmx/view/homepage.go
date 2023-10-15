package view

import (
	// "fmt"
	// "html/template"

	"net/http"

	"github.com/Censacrof/gohtmx/cmd/gohtmx/components"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	var body = components.BasePage{
		Title: "Homepage",
	}.String()

	w.Write([]byte(body))
}
