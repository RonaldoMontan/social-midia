package router

import "github.com/gorilla/mux"


//Retorna todas as rotas configuradas.
func Generate() *mux.Router {

	return mux.NewRouter()

}
