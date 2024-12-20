package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositori"
	"api/src/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



func CreatePublication(w http.ResponseWriter, r *http.Request){	

	userIdOnToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication models.Publication
	if erro = json.Unmarshal(bodyRequest, &publication); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publication.AuthorId = userIdOnToken

	if erro = publication.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositori := repositori.NewRepositoriPublication(db)
	publication.Id, erro = repositori.CreatePublication(publication)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, publication)
}

// Traz as publicações que apareceriam no feed do usuario
func SearchPublications(w http.ResponseWriter, r *http.Request){

	userIdOnToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositori := repositori.NewRepositoriPublication(db)
	publications, erro := repositori.SearchAllPublications(userIdOnToken)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, publications)

}

// Traz uma unica publicação
func SearchPublication(w http.ResponseWriter, r *http.Request){

	parameters := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(parameters["publicationsId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositori := repositori.NewRepositoriPublication(db)
	publication, erro := repositori.SearchPublicationId(publicationId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, publication)

}

func UpdatePublication(w http.ResponseWriter, r *http.Request){

	userIdOnToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(parameters["publicationsId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositori :=repositori.NewRepositoriPublication(db)
	publicationSaveOnDataBase, erro := repositori.SearchPublicationId(publicationId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicationSaveOnDataBase.AuthorId != userIdOnToken {

		response.Erro(w, http.StatusForbidden, errors.New("you do not have permission to change this"))
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication models.Publication
	if erro = json.Unmarshal(bodyRequest, &publication); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publication.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositori.AlterPublication(publicationId, publication); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)

}

func DeletePublication(w http.ResponseWriter, r *http.Request){

	userIdOnToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(parameters["publicationsId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositori :=repositori.NewRepositoriPublication(db)
	publicationSaveOnDataBase, erro := repositori.SearchPublicationId(publicationId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicationSaveOnDataBase.AuthorId != userIdOnToken {

		response.Erro(w, http.StatusForbidden, errors.New("you do not have permission to change this"))
		return
	}

	if erro = repositori.DeletePublication(publicationId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

//Traz todas as publicações do usuario
func SearchPublicationByUser(w http.ResponseWriter, r *http.Request){

	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositori := repositori.NewRepositoriPublication(db)
	publications, erro := repositori.SearchPublicationByUser(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, publications)
}

func LikePublication(w http.ResponseWriter, r *http.Request){

	parameters := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(parameters["publicationId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositori := repositori.NewRepositoriPublication(db)
	if erro = repositori.LikePublication(publicationId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
	
}