package controllers

import (
	"net/http"
	"safePasswordApi/src/database"
	login "safePasswordApi/src/models/login"
	user "safePasswordApi/src/models/user"
	"safePasswordApi/src/repository"
	"safePasswordApi/src/security/auth"
	hashEncryp "safePasswordApi/src/security/encrypt/hash"

	"github.com/labstack/echo/v4"
)

// Login
// @Summary Performs user login
// @Description Performs user login based on provided credentials
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.UserLogin true "query params"
// @Success 202 {object} login.LoginResponse "Successful login"
// @Router /Login [post]
func Login(c echo.Context) error {
	// Bind the user data from the request body
	var user user.UserLogin
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = user.Prepare()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Connect to the database
	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	// Create a user repository instance
	repo := repository.NewUserRepository(db)

	// Find the user by email in the database
	dbUser, err := repo.FindByEmail(user.Email_Hash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Compare the hashed password from the database with the provided password
	if err = hashEncryp.CompareSHA512(dbUser.Password, user.Password_Plain); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Create a login response with a JWT token
	var login login.LoginResponse
	login.Token, err = auth.CriarTokenJWT(dbUser.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Return the login response with the JWT token
	return c.JSON(http.StatusAccepted, login)
}
