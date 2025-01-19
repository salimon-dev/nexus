package main

import (
	"salimon/nexus/auth"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PUT},
	}))

	// HTTP route
	e.GET("/", rest.HeartBeatHandler)

	// register
	e.POST("/auth/register", auth.RegisterHandler)
	e.POST("/auth/register/verify", auth.VerifyRegisterHandler)

	// login
	e.POST("/auth/login", auth.LoginHandler)

	// password reset
	e.POST("/auth/password-reset", auth.PasswordResetHandler)
	e.POST("/auth/password-reset/verify", auth.VerifyPasswordResetHandler)

	e.GET("/profile", auth.GetProfileHandler, middlewares.AuthMiddleware)

	// WebSocket route
	e.GET("/sck", websocket.WsHandler)

	// E2E control Endpoints
	e.GET("/e2e/info", rest.E2EInfoHandler)
	e.POST("/e2e/reset", rest.E2EResetHandler)

	// Start the server
	port := "80"
	e.Logger.Fatal(e.Start(":" + port))
}
