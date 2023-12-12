package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func criaRoteador() http.Handler {
	roteador := chi.NewRouter()
	registraRotas(roteador)
	return roteador
}

func registraRotas(r chi.Router) {
	r.Get("/", home)
	r.Handle("/web/static/*", http.FileServer(http.FS(embeddedFiles)))
}
