package controllers

import (
	"net/http"
	"safePasswordApi/src/database"
	"safePasswordApi/src/models"
	"safePasswordApi/src/repository"
	"safePasswordApi/src/security/auth"
	hashEncryp "safePasswordApi/src/security/encrypt/hash"

	"github.com/labstack/echo/v4"
)

// Login
func Login(c echo.Context) error {
	// Bind the user data from the request body
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Connect to the database
	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	user.Email_Hash, err = hashEncryp.GenerateSHA512(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Create a user repository instance
	repo := repository.NewUserRepository(db)

	// Find the user by email in the database
	dbUser, err := repo.FindByEmail(user.Email_Hash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Compare the hashed password from the database with the provided password
	if err = hashEncryp.CompareSHA512(dbUser.Password, user.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Create a login response with a JWT token
	var login models.Login
	login.Token, err = auth.CriarTokenJWT(dbUser.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Return the login response with the JWT token
	return c.JSON(http.StatusAccepted, login)
}
