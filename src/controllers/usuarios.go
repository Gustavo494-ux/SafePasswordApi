package controllers

import "net/http"

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

// BuscarUsuario busca todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os  usuários"))
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
