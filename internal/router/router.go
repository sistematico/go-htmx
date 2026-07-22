package router

import (
	"net/http"

	"github.com/sistematico/go-htmx/internal/handlers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("GET /hello", handlers.Hello)
	mux.Handle("GET /{asset...}", assetHandler("static"))

	mux.HandleFunc("/", handlers.NotFound)

	return mux
}

// assetHandler serve os arquivos de dir sem expor o nome do diretório na URL,
// caindo no handler de 404 quando o arquivo pedido não existe.
func assetHandler(dir string) http.Handler {
	root := http.Dir(dir)
	fileServer := http.FileServer(root)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := root.Open(r.URL.Path)
		if err != nil {
			handlers.NotFound(w, r)
			return
		}
		f.Close()
		fileServer.ServeHTTP(w, r)
	})
}
