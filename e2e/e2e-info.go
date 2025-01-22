package e2e

import (
	"fmt"
	"net/http"
	"salimon/nexus/db"
	"salimon/nexus/types"

	"github.com/labstack/echo/v4"
)

type infoGetResponse struct {
	Users         []types.User         `json:"users"`
	Verifications []types.Verification `json:"verifications"`
}

func E2EInfoHandler(ctx echo.Context) error {
	response := infoGetResponse{}
	var users []types.User
	result := db.UsersModel().Select("*").Where("email ILIKE ?", "%e2e-test%").Limit(32).Find(&users)
	if result.Error != nil {
		fmt.Println(result)
		return ctx.String(http.StatusInternalServerError, "internal error")
	}
	response.Users = users

	verifications := make([]types.Verification, 128)
	index := 0
	for _, user := range users {
		var records []types.Verification
		result = db.VerificationsModel().Select("*").Where("user_id = ?", user.Id).Find(&records)
		if result.Error != nil {
			continue
		}
		for _, record := range records {
			verifications[index] = record
			index = index + 1
		}
	}

	response.Verifications = make([]types.Verification, index)
	for i := 0; i < index; i++ {
		response.Verifications[i] = verifications[i]
	}

	return ctx.JSON(http.StatusOK, response)
}
