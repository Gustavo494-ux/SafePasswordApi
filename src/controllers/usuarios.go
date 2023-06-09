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

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(c echo.Context) error {
	var usuario models.Usuario
	erro := c.Bind(&usuario)
	if erro != nil {
		return c.String(http.StatusBadRequest, erro.Error())
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioID, erro := repositorio.CriarUsuario(usuario)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	usuario, erro = repositorio.BuscarPorId(usuarioID)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	return c.JSON(http.StatusCreated, usuario)
}

// BuscarUsuarios busca um  usuário no banco de dados
func BuscarUsuario(c echo.Context) error {
	usuarioId, erro := strconv.ParseUint(c.Param("usuarioId"), 10, 64)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuario, erro := repositorio.BuscarPorId(usuarioId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	if usuario.ID == 0 {
		return c.JSON(http.StatusNotFound, errors.New("nenhum usuário foi encontrado"))
	}

	return c.JSON(http.StatusOK, usuario)
}

// BuscarUsuarios busca todos os usuários salvos no banco
func BuscarUsuarios(c echo.Context) error {
	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarios, erro := repositorio.BuscarUsuarios()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	if len(usuarios) == 0 {
		return c.JSON(http.StatusNotFound, errors.New("nenhum usuário foi encontrado"))
	}

	return c.JSON(http.StatusOK, usuarios)
}

// AtualizarUsuario Atualiza as informações de um usuário no banco
func AtualizarUsuario(c echo.Context) error {
	usuarioId, erro := strconv.ParseUint(c.Param("usuarioId"), 10, 64)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro.Error())
	}

	var usuarioRequisicao models.Usuario
	erro = c.Bind(&usuarioRequisicao)
	if erro != nil {
		return c.String(http.StatusBadRequest, erro.Error())
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorio.BuscarPorId(usuarioId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	if usuarioBanco.ID == 0 {
		return c.JSON(http.StatusNotFound, errors.New("usuário não encontrado"))
	}

	if erro = repositorio.AtualizarUsuario(usuarioId, usuarioRequisicao); erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	usuarioBanco, erro = repositorio.BuscarPorId(usuarioId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	return c.JSON(http.StatusCreated, usuarioBanco)
}

// DeletarUsuario deleta um usuário do banco de dados
func DeletarUsuario(c echo.Context) error {
	usuarioId, erro := strconv.ParseUint(c.Param("usuarioId"), 10, 64)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro.Error())
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorio.BuscarPorId(usuarioId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	if usuarioBanco.ID == 0 {
		return c.JSON(http.StatusNotFound, errors.New("usuário não encontrado"))
	}

	if erro = repositorio.DeletarUsuario(usuarioId); erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
