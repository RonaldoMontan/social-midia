package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){

	config.LoadConfig()
	fmt.Printf("API Runing in port: %d\n", config.Port)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port), r))
}