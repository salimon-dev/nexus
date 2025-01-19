package middlewares

import (
	"fmt"
	"net/http"
	"salimon/nexus/db"
	"salimon/nexus/helpers"
	"salimon/nexus/types"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorization := ctx.Request().Header.Get("Authorization")

		// Check if the header is empty or doesn't start with "Bearer "
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			return ctx.String(http.StatusUnauthorized, "unauthorized")
		}

		// Extract the token part from the header
		token := strings.TrimPrefix(authorization, "Bearer ")

		sub, err := helpers.VerifyJWT(token)

		fmt.Println(*sub)
		if err != nil || sub == nil {
			return ctx.String(http.StatusUnauthorized, "unauthorized")
		}

		var user *types.User
		result := db.UsersModel().Where("id = ?", *sub).First(&user)
		if result.Error != nil {
			fmt.Println(result.Error)
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
