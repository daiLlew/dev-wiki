package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.New("home")

		t, err := t.Parse("templates/home.html")
		if err != nil {
			w.WriteHeader(500)
			return
		}

		if err := t.Execute(w, nil); err != nil {
			w.WriteHeader(500)
		}
	}
}
