package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Representa a URL de comunicação com a API
	APIURL = ""
	//Porta onde a aplicação web está funcionando
	Port = 0
	//HashKey é utilizado para authenticar o cookie
	HashKey []byte
	//BlockKey é utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Localiza os valores no .env para carregar as variaveis de ambiente
func LoadConfig() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")

	Port, erro = strconv.Atoi(os.Getenv("WEB_APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}