package config

import (
	"log"
	"net/http"
)

func ConnectServer() error {
	var address = "localhost:8000"

	server := new(http.Server)
	server.Addr = address
	log.Println("Server is running on http://localhost:8000")

	if err := server.ListenAndServe(); err != nil {
		log.Print("Error connect to server", err)
	}

	return nil
}