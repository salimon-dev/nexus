package rest

import (
	"net/http"
	"salimon/nexus/types"

	"github.com/labstack/echo/v4"
)

type infoGetResponse struct {
	Users []types.User `json:"users"`
}

func E2EInfoHandler(ctx echo.Context) error {
	response := infoGetResponse{}
	// users, err := db.DB.Model(types.Verification{}).Select()
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }
	// response.Users = users

	// verifications := make([]types.Verification, 128)
	// response := e2eInfoResponse{}
	// users
	// user1, err := db.FindUserByEmail("user1@e2e.com")
	// fmt.Println(user1)
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }
	// if user1 != nil {
	// 	response.User1 = *user1
	// }
	// user1Verifications, err := db.GetVerificationsByUserId(user1.Id)
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }
	// if user1Verifications != nil {
	// 	response.User1Verifications = user1Verifications
	// }
	return ctx.JSON(http.StatusOK, response)
}
