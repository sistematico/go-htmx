package handlers

import (
	"net/http"

	"github.com/sistematico/go-htmx/internal/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, templates.Home, map[string]any{"Title": "go-htmx"})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.Render(w, templates.NotFound, map[string]any{"Title": "404 - go-htmx"})
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<p>Olá via htmx!</p>"))
}
