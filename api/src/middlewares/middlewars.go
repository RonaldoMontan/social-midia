package middlewares

import (
	"api/src/authentication"
	"api/src/response"
	"log"
	"net/http"
)

//Mostra infromações das rotas no terminal
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}


//Verifica se o usuario está autenticado para fazer requisições.
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){
		
		if erro := authentication.ValidateToken(r); erro != nil {
			response.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunc(w, r)
	}
}