package controllers

import (
	"api/src/authentication"
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusUnprocessableEntity)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerificarSenha(usuarioBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	var login models.Login
	login.Token, erro = authentication.CriarToken(uint64(usuarioBanco.ID))
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, login)
}
