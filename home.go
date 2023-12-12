package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	dados := make(map[string]any)
	templates.ExecuteTemplate(w, "home.html", dados)
}
