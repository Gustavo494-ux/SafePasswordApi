package repository

import (
	"errors"
	"safePasswordApi/src/models"

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
	linhasAfetadas, err := statement.RowsAffected()
	if linhasAfetadas == 0 {
		return 0, errors.New("nenhuma linha foi afetada, verifique os dados passados")
	}
	if err != nil {
		return 0, err
	}

	if erro != nil {
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

	if usuarios.ID == 0 {
		return models.Usuario{}, errors.New("nenhum usuário foi encontrado, verifique os dados passados")
	}

	if erro != nil {
		return models.Usuario{}, erro
	}
	return usuarios, nil
}

// BuscarUsuario busca todos os usuários salvos no banco
func (repositorio Usuarios) BuscarUsuarios() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	erro := repositorio.db.Select(&usuarios, "SELECT id,nome,email,senha FROM Usuarios ")
	if len(usuarios) == 0 {
		return []models.Usuario{}, errors.New("nenhum usuário foi encontrado, verifique os dados fornecidos")
	}

	if erro != nil {
		return []models.Usuario{}, erro
	}
	return usuarios, nil
}

// AtualizarUsuario Atualiza as informações de um usuário no banco
func (repositorio Usuarios) AtualizarUsuario(usuarioId uint64, usuario models.Usuario) error {
	statement, erro := repositorio.db.Exec(
		` UPDATE Usuarios SET nome =?, email =?, senha =? WHERE id =? `,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
		usuarioId,
	)
	linhasAfetadas, err := statement.RowsAffected()
	if err != nil {
		return err
	}

	if linhasAfetadas == 0 {
		return errors.New("nenhum registro foi afetado, Verifique os dados fornecidos")
	}

	if erro != nil {
		return erro
	}
	return nil
}

// BuscarPorEmail busca um usuário por email e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	usuarios := models.Usuario{}
	erro := repositorio.db.Get(&usuarios, "SELECT id,senha FROM Usuarios WHERE email = ?", email)
	if usuarios.ID == 0 {
		return models.Usuario{}, errors.New("nenhum usuário foi encontrado, verifique os dados fornecidos")
	}
	if erro != nil {
		return models.Usuario{}, erro
	}
	return usuarios, nil
}

// DeletarUsuario deleta um usuário do banco de dados
func (repositorio Usuarios) DeletarUsuario(usuarioId uint64) error {
	statement, erro := repositorio.db.Exec(
		` DELETE FROM Usuarios WHERE id =? `,
		usuarioId,
	)
	linhasAfetadas, err := statement.RowsAffected()
	if err != nil {
		return err
	}
	if linhasAfetadas == 0 {
		return errors.New("nenhum registro foi afetado, Verifique os dados fornecidos")
	}

	if erro != nil {
		return erro
	}

	return nil
}
