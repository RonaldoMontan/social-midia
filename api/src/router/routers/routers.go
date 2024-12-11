package routers

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri string
	Metodo string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

// Configure coloca todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router {

	routers := routeUsers
	routers = append(routers, routeLogin)
	routers = append(routers, routePublication...)

	for _, route := range routers {

		if route.NeedAuth {
			r.HandleFunc(route.Uri, 
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Metodo)

		} else {
			r.HandleFunc(route.Uri,
				middlewares.Logger(route.Function),
			).Methods(route.Metodo)
		}
	}

	return r
}