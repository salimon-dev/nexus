package rest

import (
	"fmt"
	"net/http"

	"salimon/nexus/db"
	"salimon/nexus/helpers"
	"salimon/nexus/middlewares"
	"salimon/nexus/types"

	"github.com/labstack/echo/v4"
)

type loginSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=5"`
}

func LoginHandler(ctx echo.Context) error {
	payload := new(loginSchema)
	if err := ctx.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	// validation errors
	vError, err := middlewares.ValidatePayload(*payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	if vError != nil {
		return ctx.JSON(http.StatusBadRequest, vError)
	}

	// fetch user based on email of verfication
	var user *types.User

	result := db.DB.Where("email = ?", payload.Email).Where("password = ?").First(user)
	if result != nil {
		fmt.Println(result.Error)
		return ctx.String(http.StatusInternalServerError, "internal error")
	}
	if user == nil {
		return ctx.String(http.StatusUnauthorized, "unauthorized")
	}

	accessToken, refreshToken, err := helpers.GenerateJWT(user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	publicUser := db.GetUserPublicObject(user)

	response := types.AuthResponse{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
		Data:         publicUser,
	}

	return ctx.JSON(http.StatusOK, response)
}
