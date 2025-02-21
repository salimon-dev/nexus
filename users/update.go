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
	Username string           `json:"username" validate:"required"`
	Password *string          `json:"password,omitempty" validate:"omitempty,gte=5"`
	Status   types.UserStatus `json:"status" validate:"required"`
	Role     types.UserRole   `json:"role" validate:"required"`
	Credit   *int32           `json:"credit" validate:"required"`
	Balance  *int32           `json:"balance" validate:"required"`
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

	record, err := db.FindUser("id = ?", id)

	if err != nil {
		fmt.Println(err)
		return helpers.InternalError(ctx)
	}

	if record == nil {
		return ctx.String(http.StatusNotFound, "not found")
	}

	if payload.Password != nil {
		passwordHash := md5.Sum([]byte(*payload.Password))
		password := hex.EncodeToString(passwordHash[:])
		record.Password = password
	}

	record.Status = payload.Status
	record.Role = payload.Role
	record.Credit = *payload.Credit
	record.Balance = *payload.Balance
	record.UpdatedAt = time.Now()

	result := db.UsersModel().Where("id = ?", id).Save(record)
	if result.Error != nil {
		fmt.Println(result.Error)
		return helpers.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, record)
}
