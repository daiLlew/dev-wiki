package main

import (
	"net/http"
	"os"

	"github.com/daiLlew/dev-wiki/config"
	"github.com/daiLlew/dev-wiki/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	cfg, err := config.Get()
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	r.HandleFunc("/wiki", handlers.HomePage(cfg))

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
