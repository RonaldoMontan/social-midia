package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//Mostra infromações das rotas no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}


//Verifica se o usuario está autenticado para fazer requisições.
func Authenticate(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){
		fmt.Println("autenticando...")
		next(w, r)
	}
}