package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositori"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositoris := repositori.NewRepositoriUsers(db)
	userSaveDb, erro := repositoris.SearchEamil(user.Email)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(userSaveDb.Password, user.Password); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authentication.CreateToken(userSaveDb.Id)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	userId := strconv.FormatUint(userSaveDb.Id, 10)

	response.JSON(w, http.StatusOK, models.DataAuthentication{
		Id:    userId,
		Token: token,
	})
}