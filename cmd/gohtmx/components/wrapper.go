package components

import (
	"bytes"
	_ "embed"
	"html/template"
)

type Wrapper struct {
}

//go:embed wrapper.html
var wrapperTemplate string

func (self Wrapper) HTML() template.HTML {
	t, err := template.New("basepage").Parse(wrapperTemplate)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	t.Execute(&buff, self)
	return template.HTML(buff.String())
}
