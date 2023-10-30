package repository

import (
	"database/sql"
	"safePasswordApi/src/logsCatalogados"
	"safePasswordApi/src/models"

	"github.com/jmoiron/sqlx"
)

type Usuario struct {
	db *sqlx.DB
}

// NovoRepositorioUsuario cria um novo repositório de usuário.
func NovoRepositorioUsuario(db *sqlx.DB) *Usuario {
	return &Usuario{db}
}

// Criar adiciona um novo usuário ao banco de dados.
func (repositorio Usuario) Criar(usuario models.Usuario) (usuarioId int64, err error) {
	statement, err := repositorio.db.Exec(
		`INSERT INTO Usuarios (nome, email,email_hash, senha) VALUES (?,?, ?,?)`,
		usuario.Nome,
		usuario.Email,
		usuario.Email_Hash,
		usuario.Senha,
	)

	if err != nil && err == sql.ErrNoRows {
		err = logsCatalogados.ErroRepositorio_DadosNaoAfetados
		return
	}

	usuarioId, err = statement.LastInsertId()
	return
}

// BuscarPorId encontra um usuário no banco de dados por ID.
func (repositorio Usuario) BuscarPorId(usuarioId uint64) (usuario models.Usuario, err error) {
	err = repositorio.db.Get(&usuario,
		`SELECT id, nome, email, criadoEm FROM Usuarios WHERE id = ?`,
		usuarioId,
	)
	if err != nil && err == sql.ErrNoRows {
		err = logsCatalogados.ErroRepositorio_DadosNaoEncontrados
	}
	return
}

// BuscarTodos recupera todos os usuários salvos no banco de dados.
func (repositorio Usuario) BuscarTodos() (usuarios []models.Usuario, err error) {
	err = repositorio.db.Select(&usuarios,
		"SELECT id, nome, email,criadoEm FROM Usuarios")
	if err != nil && err == sql.ErrNoRows {
		err = logsCatalogados.ErroRepositorio_DadosNaoEncontrados
	}
	return
}

// Atualizar atualiza as informações do usuário no banco de dados
func (repositorio Usuario) Atualizar(usuarioID uint64, usuario models.Usuario) (err error) {
	_, err = repositorio.db.Exec(
		`UPDATE Usuarios SET Nome=?, email=?, senha=? WHERE id=?`,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
		usuarioID,
	)
	if err != nil && err == sql.ErrNoRows {
		err = logsCatalogados.ErroRepositorio_DadosNaoAfetados
	}
	return
}

// BuscarPorEmail encontra um usuário por email_hash e retorna seu ID e senha com hash.
func (repositorio Usuario) BuscarPorEmail(email_hash string) (usuario models.Usuario, err error) {
	err = repositorio.db.Get(&usuario,
		"SELECT id, senha FROM Usuarios WHERE email_hash = ?",
		email_hash,
	)
	if err != nil && err == sql.ErrNoRows {
		err = logsCatalogados.ErroRepositorio_DadosNaoEncontrados
	}
	return
}

// Deletar exclui o usuário do banco de dados.
func (repositorio Usuario) Deletar(usuarioID uint64) (err error) {
	_, err = repositorio.db.Exec(
		`DELETE FROM Usuarios WHERE id = ?`,
		usuarioID,
	)
	if err != nil && err == sql.ErrNoRows {
		err = logsCatalogados.ErroRepositorio_DadosNaoAfetados
	}
	return
}
