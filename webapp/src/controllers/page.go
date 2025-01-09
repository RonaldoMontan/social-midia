package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requisitions"
	"webapp/src/response"

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
	
	publicationResponse, erro := requisitions.MakeRequisitionWithAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer publicationResponse.Body.Close()

	if publicationResponse.StatusCode >= 400 {
		response.HandleStatusCode(w, publicationResponse)
		return
	}

	var publications []models.Publication
	if erro = json.NewDecoder(publicationResponse.Body).Decode(&publications); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecTemplate(w, "home.html", struct{
		//Declaração dos campos da struct
		Publications []models.Publication
		UserId uint64
	}{//Atribuindo valores no struct
		Publications: publications,
		UserId: userId,
	})
}