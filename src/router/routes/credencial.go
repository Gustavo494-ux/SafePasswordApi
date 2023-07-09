package router

import (
	"safePasswordApi/src/controllers"
	"safePasswordApi/src/middlewares"

	"github.com/labstack/echo/v4"
)

func CredentialRoutes(e *echo.Echo) {

	// Group of routes with middleware
	credentialGroup := e.Group("/credentials")
	credentialGroup.Use(middlewares.Authenticate)

	credentialGroup.POST("", controllers.CreateCredential)
	credentialGroup.GET("", controllers.GetCredentials)
	credentialGroup.GET("/:credentialId", controllers.GetCredential)
	credentialGroup.PUT("/:credentialId", controllers.UpdateCredential)
	credentialGroup.DELETE("/:credentialId", controllers.DeleteCredential)
}
