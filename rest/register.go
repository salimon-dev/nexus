package rest

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"salimon/nexus/db"
	"salimon/nexus/mail"
	"salimon/nexus/middlewares"
	"salimon/nexus/types"

	"github.com/labstack/echo/v4"
)

type registerPayloadSchema struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=5"`
}

func RegisterHandler(ctx echo.Context) error {
	payload := new(registerPayloadSchema)
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

	userByEmail, err := db.FindUserByEmail(payload.Email)
	if err != nil {
		return err
	}
	if userByEmail != nil {
		return ctx.JSON(http.StatusBadRequest, middlewares.MakeSingleValidationError("email", "this email address is already registered"))
	}

	userByUsername, err := db.FindUserByUsername(payload.Username)
	if err != nil {
		return err
	}
	if userByUsername != nil {
		return ctx.JSON(http.StatusBadRequest, middlewares.MakeSingleValidationError("username", "this username is already registered"))
	}

	passwordHash := md5.Sum([]byte(payload.Password))
	user, err := db.InsertUser(payload.Username, payload.Email, hex.EncodeToString(passwordHash[:]), 15000, 0, types.UserRoleMember, types.UserStatusPending)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	mail.SendRegisterVerificationEmail(user)
	return ctx.String(http.StatusOK, "registered")
}
