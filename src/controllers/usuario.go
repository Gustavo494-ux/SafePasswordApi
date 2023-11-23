package controllers

import (
	"errors"
	"net/http"
	"safePasswordApi/src/database"
	enum "safePasswordApi/src/enum/geral"
	"safePasswordApi/src/logsCatalogados"
	"safePasswordApi/src/models"
	"safePasswordApi/src/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(c echo.Context) error {
	var usuario models.Usuario
	if err := c.Bind(&usuario); err != nil {
		return models.RespostaRequisicao(c).Erro(
			http.StatusBadRequest, err, logsCatalogados.ErroUsuario_JsonInvalido.Error(),
			usuario,
		).JSON()
	}

	if err := usuario.Preparar(enum.TipoPreparacao_Cadastro); err != nil {
		return models.RespostaRequisicao(c).Erro(
			http.StatusBadRequest, err, logsCatalogados.ErroUsuario_PrepararCadastro.Error(),
			usuario,
		).JSON()
	}

	db, err := database.Conectar()
	if err != nil {
		return models.RespostaRequisicao(c).Erro(
			http.StatusInternalServerError, err, logsCatalogados.LogBanco_ErroConexao,
		).JSON()
	}
	defer db.Close()

	repo := repository.NovoRepositorioUsuario(db)

	usuarioBanco, err := repo.BuscarPorEmail(usuario.Email_Hash)
	if err != logsCatalogados.ErroRepositorio_DadosNaoEncontrados {
		return models.RespostaRequisicao(c).Erro(
			http.StatusInternalServerError, errors.New(logsCatalogados.LogsUsuario_UsuarioExistente), logsCatalogados.LogsUsuario_UsuarioExistente,
		).JSON()
	}

	usuarioId, err := repo.Criar(usuario)
	if err != nil {
		return models.RespostaRequisicao(c).Erro(
			http.StatusInternalServerError, err, logsCatalogados.ErroUsuario_Cadastro.Error(),
		).JSON()
	}

	usuarioBanco, err = repo.BuscarPorId(uint64(usuarioId))
	if err != nil {
		if err == logsCatalogados.ErroRepositorio_DadosNaoEncontrados {
			return models.RespostaRequisicao(c).Erro(
				http.StatusInternalServerError, err, logsCatalogados.ErroUsuario_UsuarioNaoCadastradao.Error(),
			).JSON()
		}
		return models.RespostaRequisicao(c).Erro(
			http.StatusInternalServerError, err, logsCatalogados.ErroUsuario_GenericoConsulta.Error(),
		).JSON()
	}

	if err := usuarioBanco.Preparar(enum.TipoPreparacao_Consulta); err != nil {
		return models.RespostaRequisicao(c).Erro(
			http.StatusInternalServerError, err, logsCatalogados.ErroUsuario_PrepararConsulta.Error(),
		).JSON()
	}

	return models.RespostaRequisicao(c).Sucesso(http.StatusCreated, usuarioBanco, logsCatalogados.LogSolicitacao_AtentidaComSucesso).JSON()
}

// BuscarUsuarioPorId encontra um usuário no banco de dados por ID.
func BuscarUsuarioPorId(c echo.Context) error {
	usuarioId, err := strconv.ParseUint(c.Param("usuarioId"), 10, 64)
	if err != nil {
		return models.RespostaRequisicao(c).Erro(
			http.StatusBadRequest, err, logsCatalogados.ErroUsuario_JsonInvalido.Error(),
			c.Param("usuarioId"),
		).JSON()
	}

	db, err := database.Conectar()
	if err != nil {
		return models.RespostaRequisicao(c).Erro(
			http.StatusInternalServerError, err, logsCatalogados.LogBanco_ErroConexao,
		).JSON()
	}
	defer db.Close()

	repo := repository.NovoRepositorioUsuario(db)
	usuario, err := repo.BuscarPorId(usuarioId)
	if err != nil && err != errors.New(logsCatalogados.LogsUsuario_UsuarioNaoExistente) {
		return models.RespostaRequisicao(c).Erro(
			http.StatusInternalServerError, err, err.Error(),
		).JSON()
	}

	if usuario.ID > 0 {
		return models.RespostaRequisicao(c).Erro(
			http.StatusConflict, errors.New(logsCatalogados.LogsUsuario_UsuarioExistente), logsCatalogados.LogsUsuario_UsuarioExistente,
		).JSON()
	}

	if err := usuario.Preparar(enum.TipoPreparacao_Consulta); err != nil {
		return models.RespostaRequisicao(c).Erro(
			http.StatusInternalServerError, err, logsCatalogados.ErroUsuario_PrepararConsulta.Error(),
		).JSON()
	}

	return models.RespostaRequisicao(c).Sucesso(http.StatusCreated, usuario, logsCatalogados.LogSolicitacao_AtentidaComSucesso).JSON()
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

	if err := usuarioRequest.Preparar(enum.TipoPreparacao_Atualizar); err != nil {
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
