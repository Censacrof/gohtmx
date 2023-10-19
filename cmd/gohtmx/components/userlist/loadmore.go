package userlist

import (
	"bytes"
	_ "embed"
	"html/template"
)

type LoadMore struct {
	From int
	Size int
}

//go:embed loadmore.html
var loadMoreTemplate string

func (self LoadMore) HTML() template.HTML {
	t, err := template.New("loadmore").Parse(loadMoreTemplate)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	t.Execute(&buff, self)
	return template.HTML(buff.String())
}
