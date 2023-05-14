package controllers

import (
	"api/src/authentication"
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioID, erro := repositorio.CriarUsuario(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuario, erro = repositorio.BuscarPorId(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuario busca todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarios, erro := repositorio.BuscarUsuarios()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if len(usuarios) == 0 {
		respostas.JSON(w, http.StatusNotFound, "Nenhum Usuário foi encontrado")
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuarios busca um  usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuario, erro := repositorio.BuscarPorId(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuario.ID == 0 {
		respostas.JSON(w, http.StatusNotFound, "Nenhum Usuário foi encontrado")
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario Atualiza as informações de um usuário no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	idUsuarioLogado, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioId != idUsuarioLogado {
		respostas.JSON(w, http.StatusUnauthorized, "Você não pode alterar os dados de um usuário que não seja o seu!")
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuarioRequisicao models.Usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuarioRequisicao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorio.BuscarPorId(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuarioBanco.ID == 0 {
		respostas.JSON(w, http.StatusNotFound, "Usuário não encontrado")
	}

	if erro = repositorio.AtualizarUsuario(usuarioId, usuarioRequisicao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuarioBanco)

}

// DeletarUsuario deleta um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro!= nil {
        respostas.Erro(w, http.StatusBadRequest, erro)
        return
    }

	idUsuarioLogado, erro := authentication.ExtrairUsuarioID(r)
	if erro!= nil {
        respostas.Erro(w, http.StatusBadRequest, erro)
        return
    }

	if usuarioId!= idUsuarioLogado {
        respostas.JSON(w, http.StatusUnauthorized, "Você não pode excluir usuário que não seja o seu!")
        return
    }

	db, erro := banco.Conectar()
	if erro!= nil {
        respostas.Erro(w, http.StatusInternalServerError, erro)
        return
    }
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorio.BuscarPorId(usuarioId)
	if erro!= nil {
        respostas.Erro(w, http.StatusInternalServerError, erro)
        return
    }

	if usuarioBanco.ID == 0 {
        respostas.JSON(w, http.StatusNotFound, "Usuário não encontrado")
        return
    }

	if erro = repositorio.DeletarUsuario(usuarioId); erro!= nil {
        respostas.Erro(w, http.StatusInternalServerError, erro)
        return
    }

	respostas.JSON(w, http.StatusNoContent, nil)
}
