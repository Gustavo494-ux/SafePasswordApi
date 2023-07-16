package controllers

import (
	"errors"
	"net/http"
	"safePasswordApi/src/database"
	"safePasswordApi/src/models"
	"safePasswordApi/src/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateUser inserts a user into the database.
func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := user.Prepare("signup"); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	userId, err := repo.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user, err = repo.FindByID(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := user.Prepare("query"); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

// FindUser finds a user in the database by ID.
func FindUser(c echo.Context) error {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	user, err := repo.FindByID(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, errors.New("no user found"))
	}

	if err := user.Prepare("query"); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

// FindAllUsers retrieves all users saved in the database.
func FindAllUsers(c echo.Context) error {
	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	users, err := repo.FindAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, errors.New("no users found"))
	}

	for i := range users {
		if err = users[i].Prepare("query"); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, users)
}

// UpdateUser updates user information in the database.
func UpdateUser(c echo.Context) error {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var userRequest models.User
	if err := c.Bind(&userRequest); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	userDB, err := repo.FindByID(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if userDB.ID == 0 {
		return c.JSON(http.StatusNotFound, errors.New("user not found"))
	}

	if err := userRequest.Prepare("signup"); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := repo.UpdateUser(userId, userRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userDB, err = repo.FindByID(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := userDB.Prepare("query"); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, userDB)
}

// DeleteUser deletes a user from the database.
func DeleteUser(c echo.Context) error {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	userDB, err := repo.FindByID(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if userDB.ID == 0 {
		return c.JSON(http.StatusNotFound, errors.New("user not found"))
	}

	if err := repo.DeleteUser(userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
