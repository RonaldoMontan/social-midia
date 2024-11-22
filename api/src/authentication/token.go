package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error){

	permission := jwt.MapClaims{}

	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["userId"] = userId

	// Criar um token com método de assinatura HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)

	//Assinatura do token
	return token.SignedString([]byte("Secret"))


}