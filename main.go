package main

import (
	"salimon/nexus/auth"
	"salimon/nexus/db"
	"salimon/nexus/e2e"
	"salimon/nexus/entities"
	"salimon/nexus/invitations"
	"salimon/nexus/middlewares"
	"salimon/nexus/profile"
	"salimon/nexus/rest"
	"salimon/nexus/users"
	"salimon/nexus/websocket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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

	// login
	e.POST("/auth/login", auth.LoginHandler)
	e.POST("/auth/rotate", auth.RotateHandler)

	e.GET("/profile", profile.GetHandler, middlewares.AuthMiddleware)

	// WebSocket route
	e.GET("/sck", websocket.WsHandler)

	// -- -- Admin APIs -- --
	// invitations
	e.GET("/invitations/search", invitations.SearchHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	e.POST("/invitations/create", invitations.CreateHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	e.POST("/invitations/delete/:id", invitations.DeleteHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	e.POST("/invitations/update/:id", invitations.UpdateHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	// users
	e.GET("/users/search", users.SearchHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	e.POST("/users/create", users.CreateHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	e.POST("/users/delete/:id", users.DeleteHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	e.POST("/users/update/:id", users.UpdateHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	// entities
	e.GET("/entities/search", entities.SearchHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	e.POST("/entities/create", entities.CreateHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	e.POST("/entities/delete/:id", entities.DeleteHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)
	e.POST("/entities/update/:id", entities.UpdateHandler, middlewares.AuthMiddleware, middlewares.AdminMiddleware)

	// -- -- External APIs -- --
	// E2E control Endpoints
	e.POST("/e2e/interact", e2e.InteractHandler)

	// Start the server
	port := "80"
	e.Logger.Fatal(e.Start(":" + port))
}
