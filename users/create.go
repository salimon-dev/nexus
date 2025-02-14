package users

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"salimon/nexus/db"
	"salimon/nexus/helpers"
	"salimon/nexus/middlewares"
	"salimon/nexus/types"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type createSchema struct {
	InvitationId string           `json:"invitation_id" validate:"required,uuid"`
	Username     string           `json:"username" validate:"required"`
	Password     string           `json:"password" validate:"required,gte=5"`
	Status       types.UserStatus `json:"status" validate:"required"`
	Role         types.UserRole   `json:"role" validate:"required"`
	Credit       *int32           `json:"credit" validate:"required"`
	Balance      *int32           `json:"balance" validate:"required"`
	Usage        *int32           `json:"usage" validate:"required"`
}

func CreateHandler(ctx echo.Context) error {
	payload := new(createSchema)
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

	invitation, err := db.FindInvitation("id = ?", payload.InvitationId)
	if err != nil {
		fmt.Println(err)
		return helpers.InternalError(ctx)
	}
	if invitation == nil {
		return ctx.JSON(http.StatusBadRequest, helpers.MakeSingleValidationError("invitation_id", "invitation record not found"))
	}

	passwordHash := md5.Sum([]byte(payload.Password))
	password := hex.EncodeToString(passwordHash[:])

	record := types.User{
		Id:           uuid.New(),
		Username:     payload.Username,
		Password:     password,
		InvitationId: invitation.Id,
		Role:         payload.Role,
		Balance:      *payload.Balance,
		Credit:       *payload.Credit,
		Usage:        *payload.Usage,
		RegisteredAt: time.Now(),
		UpdatedAt:    time.Now(),
	}

	result := db.UsersModel().Create(record)
	if result.Error != nil {
		fmt.Println(result.Error)
		return helpers.InternalError(ctx)
	}
	return ctx.JSON(http.StatusOK, record)
}
