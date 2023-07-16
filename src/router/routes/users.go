package router

import (
	"safePasswordApi/src/controllers"
	"safePasswordApi/src/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	// Route without middleware
	e.POST("/user", controllers.CreateUser)

	// Group of routes with middleware
	userGroup := e.Group("/users")
	userGroup.Use(middlewares.Authenticate)

	userGroup.GET("", controllers.FindAllUsers)
	userGroup.GET("/:userId", controllers.FindUser)
	userGroup.PUT("/:userId", controllers.UpdateUser)
	userGroup.DELETE("/:userId", controllers.DeleteUser)
}
