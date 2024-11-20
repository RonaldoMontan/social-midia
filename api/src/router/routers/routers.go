package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri string
	Metodo string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

// Configute coloca todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router {

	routers := routeUsers
	routers = append(routers, routeLogin)

	for _, route := range routers {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Metodo)
	}

	return r
}