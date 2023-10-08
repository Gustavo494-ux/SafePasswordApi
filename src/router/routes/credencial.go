package router

import (
	"safePasswordApi/src/controllers"
	"safePasswordApi/src/middlewares"

	"github.com/labstack/echo/v4"
)

func CredentialRoutes(e *echo.Echo) {

	// Group of routes with middleware
	credentialGroup := e.Group("/credenciais")
	credentialGroup.Use(middlewares.Authenticate)

	credentialGroup.POST("", controllers.CriarCredencial)
	credentialGroup.GET("", controllers.BuscarCredenciais)
	credentialGroup.GET("/:credencialId", controllers.BuscarCredencialPorId)
	credentialGroup.PUT("/:credentialId", controllers.AtualizarCredencial)
	credentialGroup.DELETE("/:credentialId", controllers.DeletarCredencial)
}
