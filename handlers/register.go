package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Payload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func RegisterHandler(ctx echo.Context) error {
	payload := new(Payload)
	if err := ctx.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(payload); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, payload)

}
