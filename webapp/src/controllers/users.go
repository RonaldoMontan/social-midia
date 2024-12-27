package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/response"
)

//Chama a API para cadastrar um usuario no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	
	user, erro := json.Marshal(map[string]string{
		"name": r.FormValue("name"),
		"nick": r.FormValue("nick"),
		"email": r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	//Configura url e realiza a requisição
	url := fmt.Sprintf("%s/users", config.APIURL)
	apiResponse, erro := http.Post(url, "aplication/json", bytes.NewBuffer(user))
	
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer apiResponse.Body.Close()

	if apiResponse.StatusCode >= 400 {
		response.HandleStatusCode(w, apiResponse)
		return
	}

	response.JSON(w, apiResponse.StatusCode, nil)
}