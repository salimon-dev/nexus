package invitations

import (
	"net/http"

	"salimon/nexus/middlewares"

	"github.com/labstack/echo/v4"
)

type searchSchema struct {
	IsDepleted int `json:"is_depleted" validate:"optional,boolean"`
	PageSize   int `json:"page_size" validate:"optional,gte:6,number"`
	Page       int `json:"page" validate:"optional,number"`
}

func SearchHandler(ctx echo.Context) error {
	payload := new(searchSchema)
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

	if payload.Page == 0 {
		payload.Page = 1
	}
	if payload.PageSize == 0 {
		payload.PageSize = 10
	}

	return ctx.JSON(http.StatusOK, "ok")
}
