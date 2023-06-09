package router

import (
	"safePasswordApi/src/controllers"

	"github.com/labstack/echo/v4"
)

func RotasLogin(e *echo.Echo) {
	e.POST("/Login", controllers.Login)
}
