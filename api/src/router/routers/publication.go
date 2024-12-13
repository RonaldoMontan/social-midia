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

}