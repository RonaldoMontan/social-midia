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

//Traz as publicações dos seguidores que o usuario segue e também traz as suas publicações.
func (repositori Publications) SearchAllPublications(userIdOnToken uint64) ([]models.Publication, error){
	
	row, erro := repositori.db.Query(`
		SELECT DISTINCT 
			p.*,
			u.nick
		FROM publication p
		INNER JOIN users u on u.id = p.author_id
		INNER JOIN followers f on f.user_id  = p.author_id 
		WHERE 
			u.id = ? or f.follower_id = ?
		ORDER BY 
			1 DESC `, 
		
		userIdOnToken, userIdOnToken, )

	if erro != nil {
		return nil, erro
	}
	defer row.Close()

	var publications []models.Publication

	for row.Next() {

		var publication models.Publication

		if erro = row.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Content,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,

		); erro != nil {
			return nil, erro
		}
		publications = append(publications, publication)
	}
	return publications, nil
}

func (repositori Publications) SearchPublicationId(publicationId uint64) (models.Publication, error){
	
	row, erro := repositori.db.Query(`
		SELECT
			p.*,
			u.nick
		FROM publication p
		INNER JOIN users u ON u.id = p.author_id
		WHERE 
			p.publication_id = ?
	`, publicationId)

	if erro != nil {
		return models.Publication{}, erro
	}
	defer row.Close()

	var publication models.Publication

	if row.Next(){

		if erro = row.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Content,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		
		); erro != nil {
			return models.Publication{}, erro
		}
	}

	return publication, nil
}