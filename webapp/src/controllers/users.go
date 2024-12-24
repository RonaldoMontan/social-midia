package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		log.Fatal(erro)
	}

	response, erro := http.Post("http://localhost:5000/users", "aplication/json", bytes.NewBuffer(user))
	if erro != nil {
		log.Fatal(erro)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)
}