package repository

import (
	"api/src/models"
	"errors"

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
func (repositorio Senhas) Criar(usuarioId uint64, senha models.Senha) (uint64, error) {
	statemente, erro := repositorio.db.Exec(
		` INSERT INTO Senhas (UsuarioId,Nome,Senha) VALUES (?,?,?) `,
		usuarioId,
		senha.Nome,
		senha.Senha,
	)

	linhasAfetadas, err := statemente.RowsAffected()
	if err != nil {
		return 0, err
	}

	if linhasAfetadas == 0 {
		return 0, errors.New("nenhum registro afetado, verifique os dados fornecidos")
	}

	if erro != nil {
		return 0, erro
	}

	SenhaId, erro := statemente.LastInsertId()
	if erro != nil {
		return 0, err
	}

	return uint64(SenhaId), nil
}

// BuscarPorId retorna uma senha utilizando seu Id
func (repositorio Senhas) BuscarPorId(Id uint64) (models.Senha, error) {
	senha := models.Senha{}
	erro := repositorio.db.Get(&senha,
		` Select id, usuarioid, nome,senha,criadoem from Senhas where id = ? `,
		Id,
	)

	if senha.Id == 0 {
		return models.Senha{}, errors.New("nenhum registro retornado, verifique os dados fornecidos")
	}

	if erro != nil {
		return models.Senha{}, erro
	}
	return senha, nil
}

// BuscarPorUsuario retorna todas as senhas vinculadas a um determinado usuário
func (repositorio Senhas) BuscarPorUsuario(usuarioId uint64) ([]models.Senha, error) {
	return []models.Senha{}, nil
}

// Atualizar altera uma determinada senha utilizando seu Id como filtro
func (repositorio Senhas) Atualizar(Id uint64, Senha models.Senha) (models.Senha, error) {
	return models.Senha{}, nil
}

// Deletar exclui uma determinada senha utilizando seu Id como filtro
func (repositorio Senhas) Deletar(Id uint64) (models.Senha, error) {
	return models.Senha{}, nil
}
