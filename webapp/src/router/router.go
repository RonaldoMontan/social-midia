package route

import (
	"webapp/src/router/route"

	"github.com/gorilla/mux"
)

// Retorna todas as rotas configuradas.
func Generate() *mux.Router {

	r := mux.NewRouter()
	return router.Configure(r)

}
