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

type updateSchema struct {
	InvitationId string           `json:"invitation_id" validate:"required,uuid"`
	Username     string           `json:"username" validate:"required"`
	Password     string           `json:"password" validate:"required,gte=5"`
	Status       types.UserStatus `json:"status" validate:"required"`
	Role         types.UserRole   `json:"role" validate:"required"`
	Credit       *int32           `json:"credit" validate:"required"`
	Balance      *int32           `json:"balance" validate:"required"`
	Usage        *int32           `json:"usage" validate:"required"`
}

func UpdateHandler(ctx echo.Context) error {
	idString := ctx.Param("id")

	id, err := uuid.Parse(idString)
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusNotFound, "not found")
	}
	payload := new(updateSchema)
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

	record, err := db.FindUser("id = ?", id)

	if err != nil {
		fmt.Println(err)
		return helpers.InternalError(ctx)
	}

	if record == nil {
		return ctx.String(http.StatusNotFound, "not found")
	}

	passwordHash := md5.Sum([]byte(payload.Password))
	password := hex.EncodeToString(passwordHash[:])

	record.Status = payload.Status
	record.Role = payload.Role
	record.InvitationId = invitation.Id
	record.Usage = *payload.Usage
	record.Credit = *payload.Credit
	record.Balance = *payload.Balance
	record.InvitationId = invitation.Id
	record.Password = password
	record.UpdatedAt = time.Now()

	result := db.UsersModel().Where("id = ?", id).Save(record)
	if result.Error != nil {
		fmt.Println(result.Error)
		return helpers.InternalError(ctx)
	}
	return ctx.JSON(http.StatusOK, record)
}
