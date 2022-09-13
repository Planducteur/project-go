package main

import (
	"net/http"
	"player/handlers"

)
  func main() {
	http.HandleFunc("/ethereum/wallets/create/", handlers.CreateNewPlayer)
	http.ListenAndServe(":8090", nil)
	
  }