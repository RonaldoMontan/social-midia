package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//abstração para representar um usuario navegando na rede social
type User struct {
	Id uint64 	`json:"id, omitempty"`
	Name string `json:"name, omitempty"`
	Nick string `json:"nick, omitempty"`
	Email string `json:"email, omitempty"`
	Password string `json:"password, omitempty"`
	CreatedAt time.Time `json:"createdAt, omitempty"`
}

// chama os métodos para validar e formatar o usuario
func (user *User) Prepare(step string) error{

	if erro := user.validate(step); erro != nil {
		return erro
	}

	if erro := user.format(step); erro != nil {
		return erro
	}
	
	return nil
	
}

func (user *User) validate(step string) error {

	if user.Name == "" {
		return errors.New("The name is required")
	}

	if user.Nick == "" {
		return errors.New("The Nick is required")
	}

	if user.Email == "" {
		return errors.New("The email is required")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("The email format is invalid")
	}

	if step == "register" && user.Password == "" {
		return errors.New("The password is required")
	}

	return nil
}

func (user *User) format(step string) error {

	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {

		HashPass, erro := security.Hash(user.Password)
		if erro != nil{
			return erro
		}

		user.Password = string(HashPass)
	}
	return nil
}