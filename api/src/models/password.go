package models

// Modelo para atualização de senha
type Password struct {
	NewPassword string 	`json:"newPassword"`
	NowPassword string `json:"nowPassword"`
}