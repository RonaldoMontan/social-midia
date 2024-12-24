package router

import (
	"net/http"
	"webapp/src/controllers"
)

var routeUser = []Route{
	{
		Uri: "/create-user",
		Method: http.MethodGet,
		Function: controllers.LoadUserRegstrationPage,
		NeedAuth: false,
	},
	{
		Uri: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
}