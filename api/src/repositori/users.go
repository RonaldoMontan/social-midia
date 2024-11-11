package repositori

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

//Cria um repositório de usuários
func NewRepositoriUsers(db *sql.DB) *users {

	return &users{db}
}

//u users-> 	  Repositorio
//CreateUser-> 	  Nome do método, que recebe um modelo de usuario como parametro
//uint64, error-> retorno
func (u users) CreateUser(user models.User) (uint64, error) {

	return 0, nil
}