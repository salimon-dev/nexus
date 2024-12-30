package main

import (
	"salimon/proxy/db"
	"salimon/proxy/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.SetupDatabase()
	e := echo.New()
	e.HideBanner = true
	// Middleware
	e.Use(middleware.Recover())

	// HTTP route
	e.GET("/", handlers.HeartBeatHandler)

	// WebSocket route
	e.GET("/sck", handlers.WsHandler)

	// Start the server
	port := "80"
	e.Logger.Fatal(e.Start(":" + port))
}
