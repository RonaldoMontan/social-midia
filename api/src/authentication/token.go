package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
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

//Verifica se o token passado na requisição é valido
func ValidateToken(r  *http.Request) error {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnKeysCheck)

	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalid")

}

func extractToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnKeysCheck(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method of signature unexpected ! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}