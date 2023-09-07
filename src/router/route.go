package router

import (
	"net/http"
	_ "safePasswordApi/docs"
	router "safePasswordApi/src/router/routes"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

func Gerar() *echo.Echo {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		// c.Response().Before(routines.AnalyzingCode)
		return c.String(http.StatusOK, time.Now().Format(time.RFC3339Nano))
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	router.UserRoutes(e)
	router.RotasLogin(e)
	router.CredentialRoutes(e)

	return e
}
