package routers

import "net/http"

type Route struct {
	Uri string
	Metodo string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}