package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/models"
	"webapp/src/response"
)

// Utiliza o e-mail e senha do usuario para fazer autenticação na aplicação
func Login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	
	responseLogin, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(user))

	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer responseLogin.Body.Close()

	if responseLogin.StatusCode >= 400 {
		response.HandleStatusCode(w, responseLogin)
		return
	}

	var dataAuthentication models.DataAuthentication

	if erro = json.NewDecoder(responseLogin.Body).Decode(&dataAuthentication); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}
}