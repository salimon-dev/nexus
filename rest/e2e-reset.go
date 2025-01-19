package rest

import (
	"fmt"
	"net/http"
	"salimon/nexus/db"

	"github.com/labstack/echo/v4"
)

func E2EResetHandler(ctx echo.Context) error {
	result := db.DB.Exec("DELETE FROM verifications WHERE user_id in ( SELECT id FROM users WHERE email ILIKE ?)", "%e2e-test%")
	if result.Error != nil {
		fmt.Println(result.Error)
		return ctx.String(http.StatusInternalServerError, "interal error")
	}
	result = db.UsersModel().Where("email ILIKE ?", "%e2e-test%").Delete(nil)
	if result.Error != nil {
		fmt.Println(result.Error)
		return ctx.String(http.StatusInternalServerError, "interal error")
	}
	return ctx.String(http.StatusOK, "E2E data reset")
}
