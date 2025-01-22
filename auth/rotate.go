package auth

import (
	"fmt"
	"net/http"

	"salimon/nexus/db"
	"salimon/nexus/helpers"
	"salimon/nexus/middlewares"
	"salimon/nexus/types"

	"github.com/labstack/echo/v4"
)

type rotatePayloadSchema struct {
	Token string `json:"token" validate:"required"`
}

func RotateHandler(ctx echo.Context) error {
	payload := new(rotatePayloadSchema)
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

	claims, err := helpers.VerifyJWT(payload.Token)
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusUnauthorized, "unauthorized")
	}
	if claims == nil {
		return ctx.String(http.StatusUnauthorized, "unauthorized")
	}

	if claims.Type != "refresh" {
		return ctx.String(http.StatusUnauthorized, "unauthorized")
	}

	user, err := db.FindUser("id = ?", claims.UserID)
	if err != nil {
		fmt.Println(err)
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
