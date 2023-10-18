package userlist

import (
	"bytes"
	_ "embed"
	"html/template"
)

type LoadMoreButton struct {
	From int
	Size int
}

//go:embed loadmorebutton.html
var loadMoreButtonTemplate string

func (self LoadMoreButton) HTML() template.HTML {
	t, err := template.New("loadmorebutton").Parse(loadMoreButtonTemplate)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	t.Execute(&buff, self)
	return template.HTML(buff.String())
}
