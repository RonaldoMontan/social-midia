package routers

import (
	"api/src/router/controllers"
	"net/http"
)


var routePublication = []Route{
	
	// Criar publicação
	{
		Uri: "/publication",
		Metodo: http.MethodPost,
		Function: controllers.CreatePublication,
		NeedAuth: true,
	},
	// Busca varias publicação
	{
		Uri: "/search-publications",
		Metodo: http.MethodGet,
		Function: controllers.SearchPublications,
		NeedAuth: true,
	},
	// Busca uma publicação
	{
		Uri: "/search-publication/{publicationsId}",
		Metodo: http.MethodGet,
		Function: controllers.SearchPublication,
		NeedAuth: true,
	},
	// Atualizar uma publicação
	{
		Uri: "/update-publication/{publicationsId}",
		Metodo: http.MethodPut,
		Function: controllers.UpdatePublication,
		NeedAuth: true,
	},
	// Deletar uma publicação
	{
		Uri: "/delete-publication/{publicationsId}",
		Metodo: http.MethodDelete,
		Function: controllers.DeletePublication,
		NeedAuth: true,
	},
	// BUsca todas as publicações de um usuario
	{
		Uri: "/user/{userId}/publications",
		Metodo: http.MethodGet,
		Function: controllers.SearchPublicationByUser,
		NeedAuth: true,
	},
	// Curtir publicação
	{
		Uri: "/publication/{publicationId}/like",
		Metodo: http.MethodPost,
		Function: controllers.LikePublication,
		NeedAuth: true,
	},
}