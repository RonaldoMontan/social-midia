package security

import "golang.org/x/crypto/bcrypt"

//Recebe uma string e tranforma em hash
func Hash(password string) ([]byte, error){

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compara a senha com o hash e retorna se elas são iguais
func VerifyPassword(HashPass, StringPass string) error {

	return bcrypt.CompareHashAndPassword([]byte(HashPass), []byte(StringPass))
}