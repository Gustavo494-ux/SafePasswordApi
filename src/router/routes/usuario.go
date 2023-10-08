package router

import (
	"safePasswordApi/src/controllers"
	"safePasswordApi/src/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	// Route without middleware
	e.POST("/usuario", controllers.CriarUsuario)

	// Group of routes with middleware
	userGroup := e.Group("/usuarios")
	userGroup.Use(middlewares.Authenticate)

	userGroup.GET("", controllers.BuscarTodosUsuarios)
	userGroup.GET("/:usuarioId", controllers.BuscarUsuarioPorId)
	userGroup.PUT("/:usuarioId", controllers.AtualizarUsuario)
	userGroup.DELETE("/:usuarioId", controllers.DeletarUsuario)
}
