package routers

import "net/http"

//slice de rotas
var routeUsers = []Route{
	//Cadastra usuario
	{
		Uri: "/users",
		Metodo: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		NeedAuth: false,
	},
	//Busca todos os usuarios
	{
		Uri: "/users",
		Metodo: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		NeedAuth: false,
	},
	//Busca um usuario
	{
		Uri: "/users/{id}",
		Metodo: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		NeedAuth: false,
	},
	//Altera um usuario
	{
		Uri: "/users/{id}",
		Metodo: http.MethodPut,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		NeedAuth: false,
	},
	//Apaga um usuario
	{
		Uri: "/users/{id}",
		Metodo: http.MethodDelete,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		NeedAuth: false,
	},
}