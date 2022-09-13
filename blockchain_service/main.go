package main

import (
	"net/http"
	"blockchain_service/handlers"

)
  func main() {
	http.HandleFunc("/create/", handlers.CreateNewPlayer)
	http.ListenAndServe(":8091", nil)
	
  }