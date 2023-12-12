package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func criaRoteador() http.Handler {
	roteador := chi.NewRouter()
	registraRotas(roteador)
	return roteador
}

func registraRotas(r chi.Router) {
	r.Get("/", roteiros)
	r.Route("/roteiros", func(r chi.Router) {
		r.Get("/", roteiros)
		r.Get("/adiciona", adicionaRoteiros)
	})
	r.Handle("/web/static/*", http.FileServer(http.FS(embeddedFiles)))
	r.HandleFunc("/favicon.ico", faviconHandler)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset= UTF-8")
	data, err := embeddedFiles.ReadFile("web/static/img/favicon.ico")
	if err != nil {
		log.Println("ERRO na leitura do favicon.ico -", err)
	}
	w.Write(data)
}
