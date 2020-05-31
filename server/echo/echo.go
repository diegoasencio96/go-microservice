package server

import (
	"fmt"
	"github.com/diegoasencio96/go-microservice/config"
	status "github.com/diegoasencio96/go-microservice/controller/status"
	"github.com/diegoasencio96/go-microservice/server/general"
	"github.com/labstack/echo/v4"
	"time"
)

type EchoServer struct {
}


func (webServer EchoServer) SetupServer() {
	general.SetupLogging()

	// create web server with echo v4
	server := echo.New()
	server.Server.ReadTimeout = 15 * time.Second
	server.Server.WriteTimeout = 15 * time.Second


	// define base routes
	base := server.Group(config.Settings.BaseRoutes.BaseRoute)
	base.GET(config.Settings.BaseRoutes.StatusRoute, status.StatusController)

	// logging url server
	url := fmt.Sprintf("%s:%s", config.Settings.General.Host, config.Settings.General.Port)
	server.Logger.Fatal(server.Start(url))



}
