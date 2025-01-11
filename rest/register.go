package rest

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"salimon/nexus/db"
	"salimon/nexus/mail"
	"salimon/nexus/middlewares"
	"salimon/nexus/types"

	"github.com/google/uuid"
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

	passwordHash := md5.Sum([]byte(payload.Password))
	password := hex.EncodeToString(passwordHash[:])

	user, err := db.FindUser("username = ? AND email = ?", payload.Username, payload.Email)
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "internal error")
	}

	if user != nil {
		// user is registered and in active/inactive status
		switch user.Status {
		case types.UserStatusActive:
		case types.UserStatusInActive:
			return ctx.JSON(http.StatusBadRequest, middlewares.MakeSingleValidationError("action", "a user with same username or email already exists"))
		default:
			break
		}
		// user is pending status. so re-register
		user.Password = password
		user.RegisteredAt = time.Now()
		err := db.UpdateUser(user)
		if err != nil {
			fmt.Println(err)
			return ctx.String(http.StatusInternalServerError, "internal error")
		}
		mail.SendRegisterVerificationEmail(user)
	} else {
		user = &types.User{
			Id:           uuid.New(),
			Username:     payload.Username,
			Email:        payload.Email,
			Password:     payload.Password,
			Credit:       15000,
			Usage:        0,
			Role:         types.UserRoleMember,
			Status:       types.UserStatusPending,
			RegisteredAt: time.Now(),
			UpdatedAt:    time.Now(),
		}
		fmt.Println(user)
		err := db.InsertUser(user)
		if err != nil {
			fmt.Println(err)
			return ctx.String(http.StatusInternalServerError, "internal error")
		}
		mail.SendRegisterVerificationEmail(user)
	}
	return ctx.String(http.StatusOK, "registered")

}
