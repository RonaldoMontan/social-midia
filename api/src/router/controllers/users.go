package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositori"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request){

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()
	if erro != nil{
		log.Fatal(erro)
	}
	defer db.Close()

	repositori := repositori.NewRepositoriUsers(db)
	userId, erro := repositori.CreateUser(user)
	if erro != nil{
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID insert: %d", userId)))
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