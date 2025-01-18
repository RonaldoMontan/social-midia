package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisitions"
	"webapp/src/response"

	"github.com/gorilla/mux"
)

//Lê as informações do front (formulario) e chama a API para cadastrar uma publicação no banco de dados
func CreatePublication(w http.ResponseWriter, r*http.Request) {

	r.ParseForm()

	publication, erro := json.Marshal(map[string]string{
		"title": r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication", config.APIURL)

	responsePublication, erro := requisitions.MakeRequisitionWithAuth(r, http.MethodPost, url, bytes.NewBuffer(publication))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer responsePublication.Body.Close()

	if responsePublication.StatusCode >= 400 {
		response.HandleStatusCode(w, responsePublication)
		return
	}

	response.JSON(w, responsePublication.StatusCode, nil)
}

// Chama a API para curtir a publicação
func LikePublication(w http.ResponseWriter, r *http.Request) {

	paramters := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(paramters["publicationId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d/like", config.APIURL, publicationId)
	
	responseLike, erro := requisitions.MakeRequisitionWithAuth(r, http.MethodPost, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer responseLike.Body.Close()

	if responseLike.StatusCode >= 400 {
		response.HandleStatusCode(w, responseLike)
		return
	}

	response.JSON(w, responseLike.StatusCode, nil)
}