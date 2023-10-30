package router

import (
	"net/http"
	"safePasswordApi/src/middlewares"
	router "safePasswordApi/src/router/routes"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Gerar() *echo.Echo {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		// c.Response().Before(routines.AnalyzingCode)
		return c.String(http.StatusOK, time.Now().Format(time.RFC3339Nano))
	})

	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:        true,
		LogStatus:     true,
		LogRequestID:  true,
		LogValuesFunc: middlewares.LoggerZeroLogPersonalizado,
	}))
	e.Use(middleware.CORS())

	router.UserRoutes(e)
	router.RotasLogin(e)
	router.CredentialRoutes(e)
	return e
}
