package controllers

import (
	"net/http"
	"safePasswordApi/src/database"
	"safePasswordApi/src/models"
	"safePasswordApi/src/repository"
	"safePasswordApi/src/security"

	"github.com/labstack/echo/v4"
)

// CriarCredencial insere uma Credencial no banco de dados
func CriarCredencial(c echo.Context) error {
	var Credencial models.Credencial
	erro := c.Bind(&Credencial)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	usuarioId, erro := security.ExtrairUsuarioID(c)
	if erro != nil {
		return c.JSON(http.StatusUnauthorized, erro)
	}
	Credencial.UsuarioId = usuarioId

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorioUsuario := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorioUsuario.BuscarPorId(usuarioId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	chave, erro := usuarioBanco.GerarChaveDeCodificacaoSimetrica()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	if erro = Credencial.Preparar("salvarDados", string(chave)); erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	repositorio := repository.NovoRepositoDeCredencial(db)
	credencialID, erro := repositorio.CriarCredencial(Credencial)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	credencial, erro := repositorio.BuscarCredencialPorId(credencialID)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	return c.JSON(http.StatusCreated, credencial)
}

// BuscarCredencials busca uma Credencial no banco de dados
func BuscarCredencial(c echo.Context) error {
	return c.JSON(http.StatusOK, "Rota não implementada")
}

// BuscarCredencials busca todas as credenciais no banco de dados
func BuscarCredenciais(c echo.Context) error {
	return c.JSON(http.StatusOK, "Rota não implementada")
}

// AtualizarCredencial Atualiza as informações de uma Credencial no banco
func AtualizarCredencial(c echo.Context) error {
	return c.JSON(http.StatusNoContent, "Rota não implementada")
}

// DeletarCredencial deleta um Credencial do banco de dados
func DeletarCredencial(c echo.Context) error {
	return c.JSON(http.StatusNoContent, "Rota não implementada")
}
