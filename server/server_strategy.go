package server

import (
	echoImpl "github.com/diegoasencio96/go-microservice/server/echo"
)

type Server interface {
	SetupServer()
}

var WebServers = map[string]Server{
	"echo": echoImpl.EchoServer{},
	"gin": nil,
}
