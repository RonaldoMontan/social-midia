package repositori

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

//Traz todos os usuarios que atende ao filtro
func (repositori users) Search(nameOrNick string) ([]models.User, error){

	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	
	row, erro := repositori.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer row.Close()

	var users []models.User

	for row.Next(){

		var user models.User

		if erro = row.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}
	return users, nil
}

//Traz o usuario por Id fornecido
func (repositori users) SearchForId(userId uint64) (models.User, error){

	row, erro := repositori.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE id = ?", 
		userId,
	)

	if erro != nil {
		return models.User{}, erro
	}
	defer row.Close()

	var user models.User

	if row.Next(){

		if erro = row.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

//Atualiza as informações do usuario
func (repositori users) AlterUser(userId uint64, user models.User) error{

	statement, erro := repositori.db.Prepare(
		"update users set name = ?, nick = ?, email = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, userId); erro != nil {
		return erro
	}

	return nil
}

//Apaga as informações de um usuario
func (repositori users) DeleteUser(userId uint64) error{

	statement, erro := repositori.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userId); erro != nil {
		return erro
	}

	return nil
}

//Busca um user pelo seu email, e retorna o id com a senha hash
func (repositori users) SearchEamil(email string) (models.User, error){

	row, erro := repositori.db.Query(
		"SELECT id, password FROM users where email = ?",
		email,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if erro = row.Scan(&user.Id, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
}