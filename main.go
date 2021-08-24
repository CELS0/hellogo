package main

import (
	"github.com/CELS0/hellogo/database"
	"github.com/CELS0/hellogo/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
