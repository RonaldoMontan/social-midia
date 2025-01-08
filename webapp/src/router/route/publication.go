package router

import (
	"net/http"
	"webapp/src/controllers"
)

var routePublication = []Route {
	{
		Uri: "/publication",
		Method: http.MethodPost,
		Function: controllers.CreatePublication,
		NeedAuth: true,
	},
}