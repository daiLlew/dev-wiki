package handlers

import (
	"html/template"
	"net/http"

	"github.com/daiLlew/dev-wiki/config"
)

func HomePage(cfg *config.Config) http.HandlerFunc {
	tmpl := getTemplate()

	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "home", cfg); err != nil {
			w.WriteHeader(500)
		}
	}
}

func getTemplate() *template.Template {
	return template.Must(template.ParseFiles(
		"templates/home.html",
		"templates/nav.html",
		"templates/dev.html",
	))
}
