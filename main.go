package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Humble beginnings...")

	r := mux.NewRouter()

	r.HandleFunc("/wiki", )
}
