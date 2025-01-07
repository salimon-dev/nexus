package rest

import (
	"net/http"
	"salimon/nexus/db"
	"salimon/nexus/types"

	"github.com/labstack/echo/v4"
)

func GetProfileHandler(ctx echo.Context) error {
	user := ctx.Get("user").(*types.User)

	return ctx.JSON(http.StatusOK, db.GetUserPublicObject(user))
}
