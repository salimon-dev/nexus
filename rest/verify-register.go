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

type verifyRegisterSchema struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required"`
}

func VerifyRegisterHandler(ctx echo.Context) error {
	payload := new(verifyRegisterSchema)
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

	// get verification based on verification token
	verification, err := db.GetVerificationRecord(payload.Token)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if verification == nil {
		return ctx.String(http.StatusUnauthorized, "unauthorized")
	}

	// fetch user based on email of verfication
	var user *types.User
	result := db.UsersModel().Where("email = ?", payload.Email).First(&user)
	if err != nil {
		fmt.Println(result.Error)
		return ctx.String(http.StatusInternalServerError, "internal error")
	}
	if user == nil {
		return ctx.String(http.StatusUnauthorized, "unauthorized")
	}

	// update user status to active
	user.Status = types.UserStatusActive
	result = db.UsersModel().Save(user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return ctx.String(http.StatusInternalServerError, "internal error")
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
