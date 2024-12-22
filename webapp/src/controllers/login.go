package controllers

import "net/http"

// Renderiza a tela de login
func LoadingScreenLogin(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Tela de Login !"))
}