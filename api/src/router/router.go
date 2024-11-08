package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {

	r := mux.NewRouter()

	return routers.Configure(r)
}