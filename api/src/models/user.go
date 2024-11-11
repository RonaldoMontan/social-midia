package models

import "time"

//abstração para representar um usuario navegando na rede social
type User struct {
	Id uint64 	`json:"id,omitempty"`
	Name string `json:"name, omitempty"`
	Nick string `json:"nick, omitempty"`
	Email string `json:"email, omitempty"`
	Password string `json:"password, omitempty"`
	CreatedAt time.Time `json:"createdAt, omitempty"`

}