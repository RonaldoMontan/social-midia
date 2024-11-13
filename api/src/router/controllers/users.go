package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositori"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request){

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
	if erro != nil{
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositori := repositori.NewRepositoriUsers(db)
	user.Id, erro = repositori.CreateUser(user)
	if erro != nil{
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, user)

}

func SearchAllUsers(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Searching for all users !"))
}

func SearchUser(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Search one User !"))
}

func AlterUser(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Alter information the User !"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Delete User !"))
}