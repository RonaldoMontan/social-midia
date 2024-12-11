package repositori

import (
	"api/src/models"
	"database/sql"
)

type Publications struct{
	db *sql.DB
}

//Cria um repositorio de publicações
func NewRepositoriPublication (db *sql.DB) *Publications{
	return &Publications{db}
}

//Insere uma publicação no banco de dados
func (repositori Publications) CreatePublication(Publications models.Publication) (uint64, error){

	statement, erro := repositori.db.Prepare(`
		INSERT INTO publication (title, content, author_id) values (?,?,?)
	`)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(Publications.Title, Publications.Content, Publications.AuthorId)
	if erro != nil {
		return 0, erro
	}

	lastId, erro := result.LastInsertId()
	if erro != nil{
		return 0, erro
	}

	return uint64(lastId), nil
}