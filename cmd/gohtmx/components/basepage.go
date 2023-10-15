package components

import (
	"bytes"
	_ "embed"
	"html/template"
)

type BasePage struct {
	Title string
	Head  template.HTML
	Body  template.HTML
}

//go:embed basepage.html
var basePageTemplate string

func (self BasePage) HTML() template.HTML {
	t, err := template.New("basepage").Parse(basePageTemplate)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	t.Execute(&buff, self)
	return template.HTML(buff.String())
}
