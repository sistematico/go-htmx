package templates

import (
	"html/template"
	"net/http"
)

var (
	Home     = parse("index.html")
	NotFound = parse("404.html")
)

func parse(page string) *template.Template {
	return template.Must(template.ParseFiles("templates/base.html", "templates/"+page))
}

func Render(w http.ResponseWriter, tmpl *template.Template, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
