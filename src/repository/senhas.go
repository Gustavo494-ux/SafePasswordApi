package repository

import (
	"api/src/models"

	"github.com/jmoiron/sqlx"
)

type Senhas struct {
	db *sqlx.DB
}

// NovoRepositoDeUsuario cria um repositório de Senha
func NovoRepositoDeSenha(db *sqlx.DB) *Senhas {
	return &Senhas{db}
}

// Criar salva uma nova senha no banco de dados
func (repositorio Senhas) Criar(usuarioId uint64, senha models.Senha) (models.Senha, error) {
	return models.Senha{}, nil
}

// BuscarPorId retorna uma senha utilizando seu Id
func BuscarPorId(Id uint64) ([]models.Senha, error) {
	return []models.Senha{}, nil
}

// BuscarPorUsuario retorna todas as senhas vinculadas a um determinado usuário
func BuscarPorUsuario(usuarioId uint64) ([]models.Senha, error) {
	return []models.Senha{}, nil
}

// Atualizar altera uma determinada senha utilizando seu Id como filtro
func Atualizar(Id uint64, Senha models.Senha) (models.Senha, error) {
	return models.Senha{}, nil
}

// Deletar exclui uma determinada senha utilizando seu Id como filtro
func Deletar(Id uint64) (models.Senha, error) {
	return models.Senha{}, nil
}
