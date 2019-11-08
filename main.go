package main

import (
	"crud-golang-simple/handler"	
	"log"
	"net/http"
)

func main() {	
	Routes()
	log.Println("Server running on port :8021")
	http.ListenAndServe(":8021", nil)
}

func Routes() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)
}