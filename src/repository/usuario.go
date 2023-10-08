package repository

import (
	"database/sql"
	"errors"
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
func (repositorio Usuario) Criar(usuario models.Usuario) (uint64, error) {
	statement, err := repositorio.db.Exec(
		`INSERT INTO Usuarios (nome, email,email_hash, senha) VALUES (?,?, ?,?)`,
		usuario.Nome,
		usuario.Email,
		usuario.Email_Hash,
		usuario.Senha,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("nenhum registro afetado, verifique os dados fornecidos")
		} else {
			return 0, err
		}
	}

	usuarioID, err := statement.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(usuarioID), nil
}

// BuscarPorId encontra um usuário no banco de dados por ID.
func (repositorio Usuario) BuscarPorId(usuarioId uint64) (models.Usuario, error) {
	usuario := models.Usuario{}
	err := repositorio.db.Get(&usuario,
		`SELECT id, nome, email, criadoEm FROM Usuarios WHERE id = ?`,
		usuarioId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Usuario{}, errors.New("nenhum registro afetado, verifique os dados fornecidos")
		} else {
			return models.Usuario{}, err
		}
	}
	return usuario, nil
}

// BuscarTodos recupera todos os usuários salvos no banco de dados.
func (repositorio Usuario) BuscarTodos() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := repositorio.db.Select(&usuarios, "SELECT id, nome, email,criadoEm FROM Usuarios")
	if err != nil {
		if err == sql.ErrNoRows {
			return []models.Usuario{}, errors.New("nenhum usuário encontrado, verifique os dados fornecidos")
		}
		return []models.Usuario{}, err
	}
	return usuarios, nil
}

// Atualizar atualiza as informações do usuário no banco de dados
func (repositorio Usuario) Atualizar(usuarioID uint64, usuario models.Usuario) error {
	_, err := repositorio.db.Exec(
		`UPDATE Usuarios SET Nome=?, email=?, senha=? WHERE id=?`,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
		usuarioID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("nenhum registro afetado, verifique os dados fornecidos")
		} else {
			return err
		}
	}
	return nil
}

// BuscarPorEmail encontra um usuário por email_hash e retorna seu ID e senha com hash.
func (repositorio Usuario) BuscarPorEmail(email_hash string) (models.Usuario, error) {
	usuario := models.Usuario{}
	err := repositorio.db.Get(&usuario, "SELECT id, senha FROM Usuarios WHERE email_hash = ?", email_hash)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Usuario{}, errors.New("nenhum registro afetado, verifique os dados fornecidos")
		} else {
			return models.Usuario{}, err
		}
	}
	return usuario, nil
}

// Deletar exclui o usuário do banco de dados.
func (repositorio Usuario) Deletar(usuarioID uint64) error {
	_, err := repositorio.db.Exec(
		`DELETE FROM Usuarios WHERE id = ?`,
		usuarioID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("nenhum registro afetado, verifique os dados fornecidos")
		} else {
			return err
		}
	}
	return nil
}
