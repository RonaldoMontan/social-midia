package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Escreve informações da requisição no terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r*http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Verifica a existencia de cookies
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Read(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		nextFunction(w, r)
	}
}