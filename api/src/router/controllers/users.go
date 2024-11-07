package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Response){

	w.Write([]byte("Create User !"))
}

func SearchAllUsers(w http.ResponseWriter, r *http.Response){

	w.Write([]byte("Searching for all users !"))
}

func SearchUser(w http.ResponseWriter, r *http.Response){

	w.Write([]byte("Search one User !"))
}

func AlterUser(w http.ResponseWriter, r *http.Response){

	w.Write([]byte("Alter information the User !"))
}

func DeleteUser(w http.ResponseWriter, r *http.Response){

	w.Write([]byte("Delete User !"))
}