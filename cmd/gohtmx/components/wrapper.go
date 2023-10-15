package components

import (
	"bytes"
	_ "embed"
	"html/template"
)

type Wrapper struct {
	Content template.HTML
}

//go:embed wrapper.html
var wrapperTemplate string

func (self Wrapper) HTML() template.HTML {
	t, err := template.New("wrapper").Parse(wrapperTemplate)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	t.Execute(&buff, self)
	return template.HTML(buff.String())
}
