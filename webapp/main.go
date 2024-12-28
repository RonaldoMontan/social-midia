package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main(){

	config.LoadConfig()
	cookies.Configure()
	utils.LoadingTemplates()
	r := route.Generate()
	
	fmt.Printf("WebAPP is runing in port %d\n", config.Port)
	
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port), r ))
	
}
