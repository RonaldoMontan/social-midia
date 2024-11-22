package authentication

import (
	"api/src/config"
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

	//Assinatura do token usando uma chava do arquivo .env
	return token.SignedString([]byte(config.SecretKey))
}