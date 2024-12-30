package router

import (
	"net/http"
	"webapp/src/controllers"
)

var routeMainPage = Route{
	Uri:       "/home",
	Method:    http.MethodGet,
	Function:  controllers.LoadingMainPage,
	NeedAuth:  true,
}
