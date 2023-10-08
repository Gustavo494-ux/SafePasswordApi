package controllers

import (
	"errors"
	"net/http"
	"safePasswordApi/src/database"
	enum "safePasswordApi/src/enum/geral"
	"safePasswordApi/src/models"
	"safePasswordApi/src/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(c echo.Context) error {
	var usuario models.Usuario
	if err := c.Bind(&usuario); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := usuario.Preparar(enum.TipoPreparacao_Cadastro); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()

	repo := repository.NovoRepositorioUsuario(db)
	usuarioId, err := repo.Criar(usuario)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	usuario, err = repo.BuscarPorId(usuarioId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := usuario.Preparar(enum.TipoPreparacao_Consulta); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, usuario)
}

// BuscarUsuarioPorId encontra um usuário no banco de dados por ID.
func BuscarUsuarioPorId(c echo.Context) error {
	usuarioId, err := strconv.ParseUint(c.Param("usuarioId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repo := repository.NovoRepositorioUsuario(db)
	usuario, err := repo.BuscarPorId(usuarioId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if usuario.ID == 0 {
		return c.JSON(http.StatusNotFound, errors.New("nenhum usuário encontrado"))
	}

	if err := usuario.Preparar(enum.TipoPreparacao_Consulta); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, usuario)
}

// BuscarTodosUsuarios recupera todos os usuários salvos no banco de dados.
func BuscarTodosUsuarios(c echo.Context) error {
	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()

	repo := repository.NovoRepositorioUsuario(db)
	usuarios, err := repo.BuscarTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(usuarios) == 0 {
		return c.JSON(http.StatusNotFound, errors.New("nenhum usuário encontrado"))
	}

	for i := range usuarios {
		if err = usuarios[i].Preparar(enum.TipoPreparacao_Consulta); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, usuarios)
}

// AtualizarUsuario atualiza as informações do usuário no banco de dados.
func AtualizarUsuario(c echo.Context) error {
	usuarioId, err := strconv.ParseUint(c.Param("usuarioId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var usuarioRequest models.Usuario
	if err := c.Bind(&usuarioRequest); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repo := repository.NovoRepositorioUsuario(db)
	usuarioDB, err := repo.BuscarPorId(usuarioId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if usuarioDB.ID == 0 {
		return c.JSON(http.StatusNotFound, errors.New("usuario não encontrado"))
	}

	if err := usuarioRequest.Preparar(enum.TipoPreparacao_Consulta); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := repo.Atualizar(usuarioId, usuarioRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	usuarioDB, err = repo.BuscarPorId(usuarioId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := usuarioDB.Preparar(enum.TipoPreparacao_Consulta); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, usuarioDB)
}

// DeletarUsuario Deleta um usuário do banco de dados.
func DeletarUsuario(c echo.Context) error {
	usuarioId, err := strconv.ParseUint(c.Param("usuarioId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repo := repository.NovoRepositorioUsuario(db)
	usuarioDB, err := repo.BuscarPorId(usuarioId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if usuarioDB.ID == 0 {
		return c.JSON(http.StatusNotFound, errors.New("usuario not found"))
	}

	if err := repo.Deletar(usuarioId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
