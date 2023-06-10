package router

import (
	"safePasswordApi/src/controllers"
	"safePasswordApi/src/middlewares"

	"github.com/labstack/echo/v4"
)

func RotasCredenciais(e *echo.Echo) {

	// Grupo de rotas com middleware|
	grupoCredencial := e.Group("/credenciais")
	grupoCredencial.Use(middlewares.Autenticar)

	grupoCredencial.POST("", controllers.CriarCredencial)
	grupoCredencial.GET("", controllers.BuscarCredenciais)
	grupoCredencial.GET("/:credencialId", controllers.BuscarCredencial)
	grupoCredencial.PUT("/:credencialId", controllers.AtualizarCredencial)
	grupoCredencial.DELETE("/:credencialId", controllers.DeletarCredencial)
}
