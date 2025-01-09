package rest

import (
	"net/http"
	"salimon/nexus/db"
	"salimon/nexus/types"

	"github.com/labstack/echo/v4"
)

type e2eInfoResponse struct {
	User1              types.User           `json:"user1"`
	User1Verifications []types.Verification `json:"user1Verifications"`
}

func E2EInfoHandler(ctx echo.Context) error {
	response := e2eInfoResponse{}
	user1, err := db.FindUserByEmail("user1@e2e.com")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	response.User1 = *user1

	user1Verifications, err := db.GetVerificationsByUserId(user1.Id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	response.User1Verifications = user1Verifications

	return ctx.JSON(http.StatusOK, response)
}
