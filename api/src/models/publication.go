package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	Id uint64 	`json:"id, omitempty"`
	Title string `json:"title, omitempty"`
	Content string `json:"content, omitempty"`
	AuthorId uint64 `json:"autorId, omitempty"`
	AuthorNick string `json:"AuthorNick, omitempty"`
	Likes uint64 `json:"likes"`
	CreatedAt time.Time `json:"createdAt, omitempty"`
}

// Chama os métodos para validar e formatar a publicação
func (publication *Publication) Prepare() error {

	if erro := publication.validate(); erro != nil {
		return erro
	}

	publication.format()
	return nil
}

func (publication *Publication) validate() error {

	if publication.Title == "" {
		return errors.New("The publication must have a title")
	}

	if publication.Content == "" {
		return errors.New("The publication must have a content")
	}
	return nil
}

func (publication *Publication) format() {

	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)	
}