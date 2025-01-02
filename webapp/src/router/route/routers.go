package router

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Struct representa todas as rotas da aplicação web
type Route struct {
	Uri string
	Method string
	Function func (http.ResponseWriter, *http.Request) 
	NeedAuth bool
}

// Coloca todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router{

	routers := routeLogin
	routers = append(routers, routeUser...)
	routers = append(routers, routeMainPage)


	for _, route := range routers {

		if route.NeedAuth {
			r.HandleFunc(route.Uri,
			middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	//fixa o caminho onde está os arquivos de CSS e JS
	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}