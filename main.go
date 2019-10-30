package main

import (
	"crud-golang-simple/handler"	
	"log"
	"net/http"
)



func main() {
	// ConnectDB()
	// Routes()
	http.HandleFunc("/register", handler.Register)

	// defer db.Close()

	log.Println("Server running on port :8021")
	http.ListenAndServe(":8021", nil)
}

// func Routes() {
// }