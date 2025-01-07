package rest

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func HeartBeatHandler(c echo.Context) error {
	response := map[string]string{
		"name":        "nexus service",
		"environment": os.Getenv("ENV"),
		"time":        time.Now().Format(time.RFC3339),
	}
	return c.JSON(http.StatusOK, response)
}
