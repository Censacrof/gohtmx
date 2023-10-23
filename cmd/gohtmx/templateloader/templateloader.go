package templateloader

import "html/template"

var base_template *template.Template = nil

func initBaseTemplate() {
	base_template = template.Must(template.ParseGlob("web/template/components/*.html"))
}

func GetBaseTemplate() *template.Template {
	if base_template == nil {
		initBaseTemplate()
	}

	t, err := base_template.Clone()
	if err != nil {
		panic(err)
	}

	return t
}
