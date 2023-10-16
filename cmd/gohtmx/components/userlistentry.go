package components

import (
	"bytes"
	_ "embed"
	"html/template"
)

type UserListEntry struct {
	Avatar template.URL
}

//go:embed userlistentry.html
var UserListEntryTemplate string

func (self UserListEntry) HTML() template.HTML {
	t, err := template.New("userlistentry").Parse(UserListEntryTemplate)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	t.Execute(&buff, self)
	return template.HTML(buff.String())
}
