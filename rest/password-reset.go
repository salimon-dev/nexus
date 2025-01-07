package rest

import (
	"net/http"
	"salimon/nexus/db"
	"salimon/nexus/mail"
	"salimon/nexus/middlewares"

	"github.com/labstack/echo/v4"
)

type passwordResetPayload struct {
	Email string `json:"email" validate:"required,email"`
}

func PasswordResetHandler(ctx echo.Context) error {
	payload := new(passwordResetPayload)
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
	user, err := db.FindUserByEmail(payload.Email)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return ctx.String(http.StatusOK, "verification sent")
	}

	mail.SendPasswordResetEmail(user)
	return ctx.String(http.StatusOK, "request sent")
}
