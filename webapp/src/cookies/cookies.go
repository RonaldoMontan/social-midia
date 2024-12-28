package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)


var s *securecookie.SecureCookie


// Utiliza as variaveis de ambiente para criação do secureCookie
func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// registra as infromações de authenticação
func Save(w http.ResponseWriter, Id, token string) error {

	data := map[string]string {
		"id": Id,
		"token": token,
	}

	dataEncode, erro := s.Encode("data", data)
	if erro != nil {
		return erro
	}

	//Coloca o cookie no navegador
	http.SetCookie(w, &http.Cookie{
		Name: "data",
		Value: dataEncode,
		Path: "/",
		HttpOnly: true,
	})

	return nil
}
