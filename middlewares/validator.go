package middlewares

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type HTTPValidator struct {
	validator *validator.Validate
}

func (cv *HTTPValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errors := map[string]string{}
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, errors)
		}
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			errors[field] = "is " + tag
		}

		return echo.NewHTTPError(http.StatusBadRequest, errors)
	}
	return nil
}
