package router

import (
	"net/http"

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

	for _, route := range routers {

		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}