package handlers

import (
	"html/template"
	"net/http"
)

func HomePage() http.HandlerFunc {
	tmpl := getTemplate()

	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "home", nil); err != nil {
			w.WriteHeader(500)
		}
	}
}

func getTemplate() *template.Template {
	return template.Must(template.ParseFiles("templates/home.html", "templates/nav.html"))
}
