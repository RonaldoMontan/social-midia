package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisitions"
	
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

	url := fmt.Sprintf("%s/search-publications", config.APIURL)
	response, erro := requisitions.MakeRequisitionWithAuth(r, http.MethodGet, url, nil)

	fmt.Println(response, erro)
	utils.ExecTemplate(w, "home.html", nil)
}