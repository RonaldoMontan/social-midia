package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// Renderiza a tela de login
func LoadingScreenLogin(w http.ResponseWriter, r *http.Request){

	utils.ExecTemplate(w, "login.html", nil)
}

// Renderiza a pagina de cadastro
func LoadUserRegstrationPage(w http.ResponseWriter, r *http.Request) {

	utils.ExecTemplate(w, "register.html", nil)
}

func LoadingMainPage(w http.ResponseWriter, r *http.Request) {

	utils.ExecTemplate(w, "home.html", nil)
}