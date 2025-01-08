package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisitions"
	"webapp/src/response"
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