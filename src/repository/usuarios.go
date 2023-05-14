package repository

import (
	"api/src/models"

	"github.com/jmoiron/sqlx"
)

type Usuarios struct {
	db *sqlx.DB
}

// NovoRepositoDeUsuario cria um repositório de usuarios
func NovoRepositoDeUsuario(db *sqlx.DB) *Usuarios {
	return &Usuarios{db}
}

// BuscarPorEmail busca um usuário por email e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	usuarios := models.Usuario{}
	erro := repositorio.db.Get(&usuarios, "SELECT id,senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	return usuarios, nil
}
