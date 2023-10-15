package components

import (
	"bytes"
	_ "embed"
	"html/template"
)

type BasePage struct {
	Title string
	Head  template.HTML
}

//go:embed basepage.html
var tmpl string

func (b BasePage) String() string {
	t, err := template.New("basepage").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	t.Execute(&buff, b)
	return buff.String()
}
