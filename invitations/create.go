package invitations

import (
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
	Code           string    `json:"code" validate:"lte=12"`
	UsageRemaining int16     `json:"usage_remaining" validate:"required,gte=1"`
	ExpiresAt      time.Time `json:"expires_at" validate:"required"`
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

	user := ctx.Get("user").(*types.User)

	code := payload.Code
	if code == "" {
		code = helpers.GenerateRandomString(12)
	}

	record := types.Invitation{
		Id:             uuid.New(),
		Code:           code,
		CreatedBy:      user.Id,
		UsageRemaining: payload.UsageRemaining,
		ExpiresAt:      payload.ExpiresAt,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	result := db.InvitationsModel().Create(record)
	if result.Error != nil {
		fmt.Println(result.Error)
		return helpers.InternalError(ctx)
	}
	return ctx.JSON(http.StatusOK, record)
}
