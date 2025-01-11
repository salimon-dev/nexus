package rest

import (
	"fmt"
	"net/http"
	"salimon/nexus/db"

	"github.com/labstack/echo/v4"
)

func E2EResetHandler(ctx echo.Context) error {
	result := db.UsersModel().Where("email ILIKE ?", "%e2e-test%").Delete(nil)
	if result.Error != nil {
		fmt.Println(result.Error)
		return ctx.String(http.StatusInternalServerError, "interal error")
	}
	return ctx.String(http.StatusOK, "E2E data reset")
}
