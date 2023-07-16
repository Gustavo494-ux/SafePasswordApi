package router

import (
	router "safePasswordApi/src/router/routes"

	"github.com/labstack/echo/v4"
)

func Gerar() *echo.Echo {
	e := echo.New()
	router.UserRoutes(e)
	router.RotasLogin(e)
	router.CredentialRoutes(e)

	return e
}
