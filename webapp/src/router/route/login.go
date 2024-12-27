package router

import (
	"net/http"
	"webapp/src/controllers"
)

var routeLogin = []Route{
	{
		Uri:      "/",
		Method:   http.MethodGet,
		Function: controllers.LoadingScreenLogin,
		NeedAuth: false,
	},
	{
		Uri:      "/login",
		Method:   http.MethodGet,
		Function: controllers.LoadingScreenLogin,
		NeedAuth: false,
	},
	{
		Uri:      "/login",
		Method:   http.MethodPost,
		Function: controllers.Login,
		NeedAuth: false,
	},
}
