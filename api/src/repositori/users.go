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
func (repositori users) CreateUser(user models.User) (uint64, error) {

	statement, erro := repositori.db.Prepare(
		"INSERT INTO users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil{
		return 0, erro
	}

	lastId, erro := result.LastInsertId()
	if erro != nil{
		return 0, erro
	}

	return uint64(lastId), nil

}