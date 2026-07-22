package main

import (
	"log"
	"net/http"

	"github.com/sistematico/go-htmx/internal/router"
)

const addr = ":8080"

func main() {
	mux := router.New()

	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
