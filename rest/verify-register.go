package rest

import (
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
	user, err := db.FindUserByEmail(payload.Email)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return ctx.String(http.StatusUnauthorized, "unauthorized")
	}

	// update user status to active
	user.Status = types.UserStatusActive
	err = db.UpdateUser(user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
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
