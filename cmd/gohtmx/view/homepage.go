package view

import (
	"fmt"
	"html/template"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/homepage.html")
	if err != nil {
		message := fmt.Sprintf("%v\n", err)
		fmt.Println(message)
		http.Error(w, message, 500)
		return
	}

	t.Execute(w, struct{}{})
}
