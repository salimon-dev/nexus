package middlewares

import (
	"net/http"
	"salimon/nexus/db"
	"salimon/nexus/helpers"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorization := ctx.Request().Header.Get("authorization")

		// Check if the header is empty or doesn't start with "Bearer "
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			return ctx.String(http.StatusUnauthorized, "unauthorized")
		}

		// Extract the token part from the header
		token := strings.TrimPrefix(authorization, "Bearer ")

		sub, err := helpers.VerifyJWT(token)
		if err != nil || sub == nil {
			return ctx.String(http.StatusUnauthorized, "unauthorized")
		}

		user, err := db.FindUserById(*sub)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, "internal error")
		}
		if user == nil {
			return ctx.String(http.StatusUnauthorized, "unauthorized")
		}

		ctx.Set("sub", sub)
		ctx.Set("user", user)

		return next(ctx)
	}
}
