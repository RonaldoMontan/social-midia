package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Representa a resposta de erro na API
type ErroAPI struct {
	Erro string `json:"erro"`
}

//Retorna uma resposta em formato json para requisições
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Trata as requisições com status_code 400 ou superior
func HandleStatusCode(w http.ResponseWriter, r *http.Response) {

	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}