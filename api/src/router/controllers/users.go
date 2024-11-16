package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositori"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if erro = user.Prepare(); erro != nil{
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil{
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositoris := repositori.NewRepositoriUsers(db)
	user.Id, erro = repositoris.CreateUser(user)
	if erro != nil{
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func SearchAllUsers(w http.ResponseWriter, r *http.Request){

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))//get"user" vai pegar os parametros da url
	
	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositoris := repositori.NewRepositoriUsers(db)
	users, erro := repositoris.Search(nameOrNick)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, users)
}

func SearchUser(w http.ResponseWriter, r *http.Request){

	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
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
	user, erro := repositori.SearchForId(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, user)

}

func AlterUser(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Alter information the User !"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Delete User !"))
}