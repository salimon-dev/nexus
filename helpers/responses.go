package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InternalError(ctx echo.Context) error {
	return ctx.String(http.StatusInternalServerError, "internal error")
}

func UnauthorizedError(ctx echo.Context) error {
	return ctx.String(http.StatusUnauthorized, "unauthorized")
}
