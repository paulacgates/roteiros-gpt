package main

import "net/http"

func roteiros(w http.ResponseWriter, r *http.Request) {
	dados := make(map[string]any)
	templates.ExecuteTemplate(w, "roteiros.html", dados)
}

func adicionaRoteiros(w http.ResponseWriter, r *http.Request) {
	dados := make(map[string]any)
	templates.ExecuteTemplate(w, "adicionaRoteiros.html", dados)
}
