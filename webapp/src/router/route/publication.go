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
	{
		Uri: "/publication/{publicationId}/like",
		Method: http.MethodPost,
		Function: controllers.LikePublication,
		NeedAuth: true,
	},
	{
		Uri: "/publication/{publicationId}/unlike",
		Method: http.MethodPost,
		Function: controllers.UnlikePublication,
		NeedAuth: true,
	},
}