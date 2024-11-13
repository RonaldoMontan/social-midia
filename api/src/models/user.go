package models

import (
	"errors"
	"strings"
	"time"
)

//abstração para representar um usuario navegando na rede social
type User struct {
	Id uint64 	`json:"id,omitempty"`
	Name string `json:"name, omitempty"`
	Nick string `json:"nick, omitempty"`
	Email string `json:"email, omitempty"`
	Password string `json:"password, omitempty"`
	CreatedAt time.Time `json:"createdAt, omitempty"`
}

func (user *User) Validate() error {

	if user.Name == "" {
		return errors.New("The name is required")
	}

	if user.Nick == "" {
		return errors.New("The Nick is required")
	}

	if user.Email == "" {
		return errors.New("The email is required")
	}

	if user.Password == "" {
		return errors.New("The password is required")
	}

	return nil
}

func (user *User) Format() {

	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}