package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CriarCredencial insere uma Credencial no banco de dados
func CriarCredencial(c echo.Context) error {
	return c.JSON(http.StatusCreated, "Rota não implementada")
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
