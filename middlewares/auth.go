package middlewares

import (
	"fmt"
	"salimon/nexus/db"
	"salimon/nexus/helpers"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorization := ctx.Request().Header.Get("Authorization")

		// Check if the header is empty or doesn't start with "Bearer "
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			return helpers.UnauthorizedError(ctx)
		}

		// Extract the token part from the header
		token := strings.TrimPrefix(authorization, "Bearer ")

		claims, err := helpers.VerifyJWT(token)

		if err != nil || claims == nil {
			return helpers.UnauthorizedError(ctx)
		}

		if claims.Type != "access" {
			return helpers.UnauthorizedError(ctx)
		}

		user, err := db.FindUser("id = ?", claims.UserID)
		if err != nil {
			fmt.Println(err)
			return helpers.InternalError(ctx)
		}
		if user == nil {
			return helpers.UnauthorizedError(ctx)
		}

		ctx.Set("user", user)

		return next(ctx)
	}
}
