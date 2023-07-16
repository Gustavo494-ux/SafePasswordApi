package middlewares

import (
	"net/http"
	"safePasswordApi/src/security/auth"

	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if erro := auth.ValidarToken(c); erro != nil {
			return c.JSON(http.StatusNetworkAuthenticationRequired, "o token informado é inválido")
		}
		err := next(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return nil
	}
}
