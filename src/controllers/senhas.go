package controllers

import (
	"net/http"
)

// Criar salva uma nova senha no banco de dados
func Criar(w http.ResponseWriter, r *http.Request) {
}

// BuscarPorId retorna uma senha utilizando seu Id
func BuscarPorId(w http.ResponseWriter, r *http.Request) {
}

// BuscarPorUsuario retorna todas as senhas vinculadas a um determinado usuário
func BuscarPorUsuario(w http.ResponseWriter, r *http.Request) {
}

// Atualizar altera uma determinada senha utilizando seu Id como filtro
func Atualizar(w http.ResponseWriter, r *http.Request) {
}

// Deletar exclui uma determinada senha utilizando seu Id como filtro
func Deletar(w http.ResponseWriter, r *http.Request) {
}
