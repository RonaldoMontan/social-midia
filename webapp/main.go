package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main(){

	utils.LoadingTemplates()
	r := route.Generate()
	
	fmt.Println("WebAPP is runing in port 3000!")
	log.Fatal(http.ListenAndServe(":3000", r))
}
