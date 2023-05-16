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
)

// Criar salva uma nova senha no banco de dados
func CriarSenha(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
	}

	var senha models.Senha
	erro = json.Unmarshal(corpoRequisicao, &senha)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
	}

	if erro := senha.Validar(); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioId, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeSenha(db)
	SenhaId, erro := repositorio.Criar(usuarioId, senha)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	usuarioBanco, erro := repositorio.BuscarPorId(SenhaId)
	if erro != nil {
		respostas.Erro(w, http.StatusNotFound, erro)
	}
	respostas.JSON(w, http.StatusCreated, usuarioBanco)
}

// BuscarPorId retorna uma senha utilizando seu Id
func BuscarSenhaPorId(w http.ResponseWriter, r *http.Request) {
}

// BuscarPorUsuario retorna todas as senhas vinculadas a um determinado usuário
func BuscarSenhaPorUsuario(w http.ResponseWriter, r *http.Request) {
}

// Atualizar altera uma determinada senha utilizando seu Id como filtro
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
}

// Deletar exclui uma determinada senha utilizando seu Id como filtro
func DeletarSenha(w http.ResponseWriter, r *http.Request) {
}
