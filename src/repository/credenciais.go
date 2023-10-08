package repository

import (
	"database/sql"
	"errors"
	"safePasswordApi/src/models"

	"github.com/jmoiron/sqlx"
)

type Credencial struct {
	db *sqlx.DB
}

// NovoRepositorioCredencial cria um repositório de credenciais
func NovoRepositorioCredencial(db *sqlx.DB) *Credencial {
	return &Credencial{db}
}

// Criar insere uma nova credencial
func (repo *Credencial) Criar(credential models.Credencial) (uint64, error) {
	resultado, err := repo.db.Exec(
		`INSERT INTO Credenciais (usuarioId, descricao, siteUrl, login, Senha) VALUES (?, ?, ?, ?, ?)`,
		credential.UsuarioId,
		credential.Descricao,
		credential.SiteUrl,
		credential.Login,
		credential.Senha,
	)
	if err != nil {
		return 0, err
	}

	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		return 0, err
	}
	if linhasAfetadas == 0 {
		return 0, errors.New("nenhuma linha afetada, verifique os dados fornecidos")
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarPorId busca uma credencial por seu ID
func (repo *Credencial) BuscarPorId(credentialID uint64) (models.Credencial, error) {
	credential := models.Credencial{}
	err := repo.db.Get(&credential,
		`SELECT id, usuarioId, descricao, siteUrl, login, senha, criadoem FROM Credenciais WHERE id = ?`,
		credentialID,
	)
	if err == sql.ErrNoRows {
		return models.Credencial{}, errors.New("nenhuma credencial encontrada, verifique os dados fornecidos")
	}
	if err != nil {
		return models.Credencial{}, err
	}

	return credential, nil
}

// BuscarTodos busca todas as credenciais armazenadas no banco de dados para um determinado usuário
func (repo *Credencial) BuscarTodos(userID uint64) ([]models.Credencial, error) {
	credentials := []models.Credencial{}
	err := repo.db.Select(&credentials,
		"SELECT id, usuarioId, descricao, siteUrl, login, senha, criadoem FROM Credenciais WHERE UsuarioId = ?",
		userID,
	)
	if len(credentials) == 0 {
		return []models.Credencial{}, errors.New("nenhuma credencial encontrada, verifique os dados fornecidos")
	}
	if err != nil {
		return []models.Credencial{}, err
	}

	return credentials, nil
}

// Atualizar atualiza as informações de uma credencial no banco de dados
func (repo *Credencial) Atualizar(credentialID uint64, credential models.Credencial) error {
	resultado, err := repo.db.Exec(
		`UPDATE Credenciais SET descricao=?, siteUrl=?, login=?, senha=? WHERE id=?`,
		credential.Descricao,
		credential.SiteUrl,
		credential.Login,
		credential.Senha,
		credentialID,
	)
	if err != nil {
		return err
	}

	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		return err
	}
	if linhasAfetadas == 0 {
		return errors.New("nenhuma linha afetada, verifique os dados fornecidos")
	}

	return nil
}

// Deletar deleta uma credencial do banco de dados.
func (repo *Credencial) Deletar(credentialID uint64) error {
	resultado, err := repo.db.Exec(
		`DELETE FROM Credenciais WHERE id = ?`,
		credentialID,
	)
	if err != nil {
		return err
	}

	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		return err
	}
	if linhasAfetadas == 0 {
		return errors.New("nenhuma linha afetada, verifique os dados fornecidos")
	}

	return nil
}
