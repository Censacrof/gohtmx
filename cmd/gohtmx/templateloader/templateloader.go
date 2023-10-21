package templateloader

import (
	"fmt"
	"html/template"
)

var cached_template *template.Template = nil

func initTemplate() {
	cached_template = template.Must(template.ParseFiles(
		"web/template/basepage.comp.html",
	))

	cached_template = template.Must(cached_template.ParseFiles(
		"web/template/homepage.view.html",
	))

	for _, t := range cached_template.Templates() {
		fmt.Println(t.Name())
	}
}

func GetTemplate() *template.Template {
	if cached_template == nil {
		initTemplate()
	}

	return cached_template
}
