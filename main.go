package main

import "github.com/CELS0/hellogo/server"

func main() {
	server := server.NewServer()

	server.Run()
}
