package main

import (
	"hygya-api/database"
	"hygya-api/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()
}
