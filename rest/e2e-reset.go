package rest

import (
	"net/http"
	"salimon/nexus/db"

	"github.com/labstack/echo/v4"
)

func E2EResetHandler(ctx echo.Context) error {
	err := db.DeleteE2ETestUsers()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, "E2E data reset")
}
