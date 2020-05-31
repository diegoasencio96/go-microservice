package main

import (
	"github.com/diegoasencio96/go-microservice/server"
	"os"
)

func main() {
	webServer := os.Getenv("WEB_SERVER")

	server, exists := server.WebServers[webServer]
	if exists {
		server.SetupServer()
	}

}
