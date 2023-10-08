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
	var usuario models.Usuario
	err := c.Bind(&usuario)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	usuario.Email_Hash, err = hashEncryp.GenerateSHA512(usuario.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	repo := repository.NovoRepositorioUsuario(db)

	dbusuario, err := repo.BuscarPorEmail(usuario.Email_Hash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err = hashEncryp.CompareSHA512(dbusuario.Senha, usuario.Senha); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var login models.Login
	login.Token, err = auth.CriarTokenJWT(dbusuario.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusAccepted, login)
}
