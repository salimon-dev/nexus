package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type HttpValidator struct {
	validator *validator.Validate
}

func (cv *HttpValidator) Validate(i interface{}) error {
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

type EventResponse struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"Body"`
}

func SuccessResponse(c echo.Context) error {
	result := map[string]interface{}{"statusCode": 200, "body": "success"}
	return c.JSON(http.StatusOK, result)
}

func FailureResponse(c echo.Context) error {
	result := map[string]interface{}{"statusCode": 503, "body": "internal error"}
	return c.JSON(http.StatusOK, result)
}
