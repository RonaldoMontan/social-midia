package models

import "time"

type Publication struct {
	Id uint64 	`json:"id, omitempty"`
	Title string `json:"title, omitempty"`
	Content string `json:"content, omitempty"`
	AuthorId uint64 `json:"autorId, omitempty"`
	AuthorNick string `json:"AuthorNick, omitempty"`
	Likes uint64 `json:"likes"`
	CreatedAt time.Time `json:"createdAt, omitempty"`
}