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
		NeedAuth: true,
	},
	//Busca um usuario
	{
		Uri: "/users/{id}",
		Metodo: http.MethodGet,
		Function: controllers.SearchUser,
		NeedAuth: true,
	},
	//Altera um usuario
	{
		Uri: "/users/{id}",
		Metodo: http.MethodPut,
		Function: controllers.AlterUser,
		NeedAuth: true,
	},
	//Apaga um usuario
	{
		Uri: "/users/{id}",
		Metodo: http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: true,
	},
	//Seguir um usuario
	{
		Uri: "/users/{id}/follow",
		Metodo: http.MethodPost,
		Function: controllers.FollowUser,
		NeedAuth: true,
	},
	//Deixar de seguir
	{
		Uri: "/users/{id}/unfollow",
		Metodo: http.MethodPost,
		Function: controllers.UnfollowUser,
		NeedAuth: true,
	},
	//Busca seguidores.
	{
		Uri: "/users/{id}/followers",
		Metodo: http.MethodGet,
		Function: controllers.SearchFollowers,
		NeedAuth: true,
	},
	//quem estou seguindo
	{
		Uri: "/users/{id}/following",
		Metodo: http.MethodGet,
		Function: controllers.SearchFollowing,
		NeedAuth: true,
	},
	//atualizar senha
	{
		Uri: "/users/{id}/update-password",
		Metodo: http.MethodPost,
		Function: controllers.UpdatePassword,
		NeedAuth: true,
	},
}