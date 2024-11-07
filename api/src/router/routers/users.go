package routers

import (
	"api/src/router/controllers"
	"net/http"
)

//slice de rotas
var routeUsers = []Route{
	//Cadastra usuario
	{
		Uri: "/users",
		Metodo: http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	//Busca todos os usuarios
	{
		Uri: "/users",
		Metodo: http.MethodGet,
		Function: controllers.SearchAllUsers,
		NeedAuth: false,
	},
	//Busca um usuario
	{
		Uri: "/users/{id}",
		Metodo: http.MethodGet,
		Function: controllers.SearchUser,
		NeedAuth: false,
	},
	//Altera um usuario
	{
		Uri: "/users/{id}",
		Metodo: http.MethodPut,
		Function: controllers.AlterUser,
		NeedAuth: false,
	},
	//Apaga um usuario
	{
		Uri: "/users/{id}",
		Metodo: http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: false,
	},
}