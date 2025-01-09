package main

import (
	"salimon/nexus/db"
	"salimon/nexus/mail"
	"salimon/nexus/middlewares"
	"salimon/nexus/rest"
	"salimon/nexus/websocket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	mail.SetupEmailQueue()
	db.SetupDatabase()
	e := echo.New()
	e.HideBanner = true
	// Middleware
	e.Use(middleware.Recover())

	// HTTP route
	e.GET("/", rest.HeartBeatHandler)

	// register
	e.POST("/auth/register", rest.RegisterHandler)
	e.POST("/auth/register/verify", rest.VerifyRegisterHandler)

	// login
	e.POST("/auth/login", rest.LoginHandler)

	// password reset
	e.POST("/auth/password-reset", rest.PasswordResetHandler)
	e.POST("/auth/password-reset/verify", rest.VerifyPasswordResetHandler)

	e.GET("/profile", rest.GetProfileHandler, middlewares.AuthMiddleware)

	// WebSocket route
	e.GET("/sck", websocket.WsHandler)

	// E2E control Endpoints
	e.GET("/e2e/info", rest.E2EInfoHandler)
	e.POST("/e2e/reset", rest.E2EResetHandler)

	// Start the server
	port := "80"
	e.Logger.Fatal(e.Start(":" + port))
}
