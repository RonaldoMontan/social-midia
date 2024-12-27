package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	token, _ := io.ReadAll(responseLogin.Body)

	fmt.Println(responseLogin.StatusCode, string(token))
}