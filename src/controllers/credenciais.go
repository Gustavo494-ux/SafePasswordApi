package controllers

import (
	"errors"
	"net/http"
	"safePasswordApi/src/database"
	"safePasswordApi/src/models"
	"safePasswordApi/src/repository"
	"safePasswordApi/src/security/auth"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateCredential inserts a Credencial into the database
func CreateCredential(c echo.Context) error {
	var credential models.Credencial
	err := c.Bind(&credential)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID, err := auth.ExtrairUsuarioID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}
	credential.UsuarioId = userID

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()

	if err = credential.Prepare("saveData"); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	credRepo := repository.NewCredentialRepository(db)
	credID, err := credRepo.CreateCredential(credential)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	credDB, err := credRepo.FindCredentialByID(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, credDB)
}

// GetCredential retrieves a Credencial from the database
func GetCredential(c echo.Context) error {
	credID, err := strconv.ParseUint(c.Param("credentialID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	credRepo := repository.NewCredentialRepository(db)
	cred, err := credRepo.FindCredentialByID(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if cred.Id == 0 {
		return c.JSON(http.StatusNotFound, errors.New("no credential found"))
	}

	if err = cred.Prepare("retrieveData"); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, cred)
}

// GetCredentials retrieves all credentials of the logged-in user
func GetCredentials(c echo.Context) error {
	userID, err := auth.ExtrairUsuarioID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	credRepo := repository.NewCredentialRepository(db)
	credsDB, err := credRepo.FindCredentials(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if len(credsDB) == 0 {
		return c.JSON(http.StatusNotFound, errors.New("no credentials found"))
	}

	var decryptedCredentials []models.Credencial
	for _, cred := range credsDB {
		if err := cred.Prepare("retrieveData"); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		decryptedCredentials = append(decryptedCredentials, cred)
	}

	return c.JSON(http.StatusOK, decryptedCredentials)
}

// UpdateCredential updates the information of a Credencial in the database
func UpdateCredential(c echo.Context) error {
	credID, err := strconv.ParseUint(c.Param("credentialID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var requestCredential models.Credencial
	err = c.Bind(&requestCredential)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID, err := auth.ExtrairUsuarioID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	credRepo := repository.NewCredentialRepository(db)
	credDB, err := credRepo.FindCredentialByID(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if credDB.Id == 0 {
		return c.JSON(http.StatusNotFound, errors.New("no credential found"))
	}

	if credDB.UsuarioId != userID {
		return c.JSON(http.StatusNotFound, errors.New("cannot update a credential of another user"))
	}

	requestCredential.UsuarioId = credDB.UsuarioId

	if err = requestCredential.Prepare("saveData"); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = credRepo.UpdateCredential(credDB.Id, requestCredential)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	credDB, err = credRepo.FindCredentialByID(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err = credDB.Prepare("retrieveData"); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, credDB)
}

// DeleteCredential deletes a Credencial from the database
func DeleteCredential(c echo.Context) error {
	credID, err := strconv.ParseUint(c.Param("credentialID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID, err := auth.ExtrairUsuarioID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	credRepo := repository.NewCredentialRepository(db)
	credDB, err := credRepo.FindCredentialByID(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if credDB.UsuarioId != userID {
		return c.JSON(http.StatusNotFound, errors.New("cannot delete a credential of another user"))
	}

	if err := credRepo.DeleteCredential(credID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, "Credential deleted successfully!")
}
