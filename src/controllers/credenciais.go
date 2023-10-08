package controllers

import (
	"errors"
	"net/http"
	"safePasswordApi/src/database"
	enum "safePasswordApi/src/enum/geral"
	"safePasswordApi/src/models"
	"safePasswordApi/src/repository"
	"safePasswordApi/src/security/auth"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CriarCredencial insere uma Credencial no banco de dados
func CriarCredencial(c echo.Context) error {
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

	if err = credential.Preparar(enum.TipoPreparacao_Cadastro); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	credRepo := repository.NovoRepositorioCredencial(db)
	credID, err := credRepo.Criar(credential)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	credDB, err := credRepo.BuscarPorId(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err = credDB.Preparar(enum.TipoPreparacao_Consulta); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, credDB)
}

// BuscarCredencialPorId recupera uma credencial do banco de dados utilizando seu id
func BuscarCredencialPorId(c echo.Context) error {
	credID, err := strconv.ParseUint(c.Param("credencialId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	credRepo := repository.NovoRepositorioCredencial(db)
	cred, err := credRepo.BuscarPorId(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if cred.Id == 0 {
		return c.JSON(http.StatusNotFound, errors.New("credencial não encontrada"))
	}

	if err = cred.Preparar(enum.TipoPreparacao_Consulta); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, cred)
}

// BuscarCredenciais recupera todas as credenciais do usuário logado
func BuscarCredenciais(c echo.Context) error {
	userID, err := auth.ExtrairUsuarioID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	credRepo := repository.NovoRepositorioCredencial(db)
	credsDB, err := credRepo.BuscarTodos(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if len(credsDB) == 0 {
		return c.JSON(http.StatusNotFound, errors.New("nenhuma credencial encontrada"))
	}

	var decryptedCredentials []models.Credencial
	for _, cred := range credsDB {
		if err := cred.Preparar(enum.TipoPreparacao_Consulta); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		decryptedCredentials = append(decryptedCredentials, cred)
	}

	return c.JSON(http.StatusOK, decryptedCredentials)
}

// AtualizarCredencial atualiza as informações de uma Credencial no banco de dados
func AtualizarCredencial(c echo.Context) error {
	credID, err := strconv.ParseUint(c.Param("credentialId"), 10, 64)
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

	credRepo := repository.NovoRepositorioCredencial(db)
	credDB, err := credRepo.BuscarPorId(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if credDB.Id == 0 {
		return c.JSON(http.StatusNotFound, errors.New("nenhuma credencial encontrada"))
	}

	if credDB.UsuarioId != userID {
		return c.JSON(http.StatusNotFound, errors.New("não é possível atualizar uma credencial de outro usuário"))
	}

	requestCredential.UsuarioId = credDB.UsuarioId

	if err = requestCredential.Preparar(enum.TipoPreparacao_Atualizar); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = credRepo.Atualizar(credDB.Id, requestCredential)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	credDB, err = credRepo.BuscarPorId(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err = credDB.Preparar(enum.TipoPreparacao_Consulta); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, credDB)
}

// DeletarCredencial exclui uma credencial do banco de dados
func DeletarCredencial(c echo.Context) error {
	credID, err := strconv.ParseUint(c.Param("credentialId"), 10, 64)
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

	credRepo := repository.NovoRepositorioCredencial(db)
	credDB, err := credRepo.BuscarPorId(credID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if credDB.UsuarioId != userID {
		return c.JSON(http.StatusNotFound, errors.New("não é possível excluir uma credencial de outro usuário"))
	}

	if err := credRepo.Deletar(credID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, "credencial deletada com sucesso!")
}
