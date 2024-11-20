package routers

import (
	"api/src/router/controllers"
	"net/http"
)

var routeLogin = Route{
	Uri: "/login",
	Metodo: http.MethodPost,
	Function: controllers.Login,
	NeedAuth: false,
}