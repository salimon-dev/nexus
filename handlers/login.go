package handlers

import (
	"net/http"

	"salimon/proxy/db"
	"salimon/proxy/helpers"
	"salimon/proxy/middlewares"
	"salimon/proxy/types"

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
	user, err := db.FindUserByAuth(payload.Email, payload.Password)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
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
