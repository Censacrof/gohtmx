package components

import (
	"bytes"
	_ "embed"
	"html/template"
)

type UserList struct {
}

//go:embed userlist.html
var userListTemplate string

func (self UserList) HTML() template.HTML {
	t, err := template.New("userlist").Parse(userListTemplate)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	t.Execute(&buff, self)
	return template.HTML(buff.String())
}
