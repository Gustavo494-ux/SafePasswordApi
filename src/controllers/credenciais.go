package controllers

import (
	"errors"
	"net/http"
	"safePasswordApi/src/database"
	"safePasswordApi/src/models"
	"safePasswordApi/src/repository"
	"safePasswordApi/src/security"
	"strconv"

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
	credencialId, erro := strconv.ParseUint(c.Param("credencialId"), 10, 64)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeCredencial(db)
	credencial, erro := repositorio.BuscarCredencialPorId(credencialId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	if credencial.Id == 0 {
		return c.JSON(http.StatusNotFound, errors.New("nenhuma credencial foi encontrado"))
	}

	repositorioUsuario := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorioUsuario.BuscarPorId(credencial.UsuarioId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	chave, erro := usuarioBanco.GerarChaveDeCodificacaoSimetrica()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	if erro = credencial.Preparar("consultarDados", string(chave)); erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	return c.JSON(http.StatusOK, credencial)
}

// BuscarCredencials busca todas as credenciais do usuário logado
func BuscarCredenciais(c echo.Context) error {
	usuarioId, erro := security.ExtrairUsuarioID(c)
	if erro != nil {
		return c.JSON(http.StatusUnauthorized, erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeCredencial(db)
	credenciaisBanco, erro := repositorio.BuscarCredenciais(usuarioId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	if len(credenciaisBanco) == 0 {
		return c.JSON(http.StatusNotFound, errors.New("nenhuma credencial foi encontrado"))
	}

	repositorioUsuario := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorioUsuario.BuscarPorId(usuarioId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	chave, erro := usuarioBanco.GerarChaveDeCodificacaoSimetrica()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	var credenciaisDescriptografadas []models.Credencial
	for _, credencial := range credenciaisBanco {
		if erro := credencial.Preparar("consultarDados", string(chave)); erro != nil {
			return c.JSON(http.StatusBadRequest, erro)
		}
		credenciaisDescriptografadas = append(credenciaisDescriptografadas, credencial)
	}

	return c.JSON(http.StatusOK, credenciaisDescriptografadas)
}

// AtualizarCredencial Atualiza as informações de uma Credencial no banco
func AtualizarCredencial(c echo.Context) error {
	credencialId, erro := strconv.ParseUint(c.Param("credencialId"), 10, 64)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	var credencialRequisicao models.Credencial
	erro = c.Bind(&credencialRequisicao)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	usuarioId, erro := security.ExtrairUsuarioID(c)
	if erro != nil {
		return c.JSON(http.StatusUnauthorized, erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}
	defer db.Close()

	repositorioUsuario := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorioUsuario.BuscarPorId(usuarioId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	chave, erro := usuarioBanco.GerarChaveDeCodificacaoSimetrica()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	repositorio := repository.NovoRepositoDeCredencial(db)
	credencialBanco, erro := repositorio.BuscarCredencialPorId(credencialId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	if credencialBanco.Id == 0 {
		return c.JSON(http.StatusNotFound, errors.New("nenhuma credencial foi encontrado"))
	}

	if credencialBanco.UsuarioId != usuarioId {
		return c.JSON(http.StatusNotFound, errors.New("não é possível atualizar uma credencial de outro usuário"))
	}

	credencialRequisicao.UsuarioId = credencialBanco.UsuarioId

	if erro = credencialRequisicao.Preparar("salvarDados", string(chave)); erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	erro = repositorio.AtualizarCredencial(credencialBanco.Id, credencialRequisicao)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	credencialBanco, erro = repositorio.BuscarCredencialPorId(credencialId)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	if erro = credencialBanco.Preparar("consultarDados", string(chave)); erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	return c.JSON(http.StatusOK, credencialBanco)
}

// DeletarCredencial deleta um Credencial do banco de dados
func DeletarCredencial(c echo.Context) error {
	return c.JSON(http.StatusNoContent, "Rota não implementada")
}
