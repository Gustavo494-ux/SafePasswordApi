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

// CriarUsuario Adiciona um novo usuário
func (repositorio Usuarios) CriarUsuario(Usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Exec(
		` INSERT INTO Usuarios (nome, email, senha ) values (?,?,?) `,
		Usuario.Nome,
		Usuario.Email,
		Usuario.Senha,
	)
	if erro != nil {
		return 0, erro
	}

	linhasAfetadas, erro := statement.RowsAffected()
	if erro != nil || linhasAfetadas == 0 {
		return 0, erro
	}

	usuarioID, erro := statement.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(usuarioID), nil
}

// BuscarPorId busca um usuário pelo ID
func (repositorio Usuarios) BuscarPorId(usuarioId uint64) (models.Usuario, error) {
	usuarios := models.Usuario{}
	erro := repositorio.db.Get(&usuarios,
		` SELECT id,nome,email,criadoem FROM Usuarios WHERE id = ? `,
		usuarioId,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}
	return usuarios, nil
}

// BuscarPorEmail busca um usuário por email e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	usuarios := models.Usuario{}
	erro := repositorio.db.Get(&usuarios, "SELECT id,senha FROM Usuarios WHERE email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	return usuarios, nil
}
