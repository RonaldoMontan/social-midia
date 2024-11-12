package response

import (
	"encoding/json"
	"log"
	"net/http"
)

//Padrão de resposta em json
func JSON(w http.ResponseWriter, statusCode int, dados interface{}){

	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil{
		log.Fatal(erro)
	}

}

//Retorna um erro no formato json
func Erro(w http.ResponseWriter, statusCode int, erro error){

	JSON(w, statusCode, struct{
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}