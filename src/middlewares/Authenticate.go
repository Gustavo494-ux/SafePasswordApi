package middlewares

import (
	"fmt"
	"net/http"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/security/auth"

	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if erro := auth.ValidarToken(c); erro != nil {
			logger.Logger().Info(fmt.Sprintf("Token %s inválido", auth.ExtrairToken(c)))
			return c.JSON(http.StatusNetworkAuthenticationRequired, "o token informado é inválido")
		}
		err := next(c)
		if err != nil {
			logger.Logger().Error(fmt.Sprintf("Ocorreu um erro na requisição, token: %s", auth.ExtrairToken(c)), err)
			return c.JSON(http.StatusUnauthorized, err)
		}
		return nil
	}
}
