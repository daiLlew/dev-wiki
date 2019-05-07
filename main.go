package main

import (
	"net/http"

	"github.com/daiLlew/dev-wiki/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomePage())
	r.HandleFunc("/wiki", handlers.HomePage())

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
