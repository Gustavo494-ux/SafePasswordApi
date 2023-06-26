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
	erro := c.Bind(&usuario)
	if erro != nil {
		return c.String(http.StatusBadRequest, erro.Error())
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	if erro = hashEncryp.CompareSHA512(usuarioBanco.Senha, usuario.Senha); erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	var login models.Login
	login.Token, erro = auth.CriarTokenJWT(usuarioBanco.ID)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	return c.JSON(http.StatusOK, login)
}
