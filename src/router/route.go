package router

import (
	"net/http"
	router "safePasswordApi/src/router/routes"
	"safePasswordApi/src/routines"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Gerar() *echo.Echo {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		c.Response().Before(routines.AnalyzingCode)
		return c.String(http.StatusOK, time.Now().Format(time.RFC3339Nano))
	})

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	router.UserRoutes(e)
	router.RotasLogin(e)
	router.CredentialRoutes(e)

	return e
}
