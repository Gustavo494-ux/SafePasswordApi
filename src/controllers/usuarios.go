package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	w.Write([]byte("Buscando um usuário"))
}

// AtualizarUsuario Atualiza as informações de um usuário no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário"))
}

// DeletarUsuario deleta um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}
