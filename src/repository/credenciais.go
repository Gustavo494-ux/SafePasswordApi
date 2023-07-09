package repository

import (
	"database/sql"
	"errors"
	"safePasswordApi/src/models"

	"github.com/jmoiron/sqlx"
)

type Credential struct {
	db *sqlx.DB
}

// NewCredentialRepository create a credential repository
func NewCredentialRepository(db *sqlx.DB) *Credential {
	return &Credential{db}
}

// CreateCredential insert a new credential
func (repo *Credential) CreateCredential(credential models.Credencial) (uint64, error) {
	resultado, err := repo.db.Exec(
		`INSERT INTO Credenciais (UsuarioId, Descricao, SiteUrl, Login, Senha) VALUES (?, ?, ?, ?, ?)`,
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
		return 0, errors.New("no rows affected, check the provided data")
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// FindCredentialByID fetch a credential by its ID
func (repo *Credential) FindCredentialByID(credentialID uint64) (models.Credencial, error) {
	credential := models.Credencial{}
	err := repo.db.Get(&credential,
		`SELECT id, UsuarioId, Descricao, SiteUrl, Login, Senha, CriadoEm FROM Credenciais WHERE id = ?`,
		credentialID,
	)
	if err == sql.ErrNoRows {
		return models.Credencial{}, errors.New("no credential found, check the provided data")
	}
	if err != nil {
		return models.Credencial{}, err
	}

	return credential, nil
}

// FindCredentials fetches all credentials stored in the database for a given user
func (repo *Credential) FindCredentials(userID uint64) ([]models.Credencial, error) {
	credentials := []models.Credencial{}
	err := repo.db.Select(&credentials,
		"SELECT id, UsuarioId, Descricao, SiteUrl, Login, Senha, CriadoEm FROM Credenciais WHERE UsuarioId = ?",
		userID,
	)
	if len(credentials) == 0 {
		return []models.Credencial{}, errors.New("no credentials found, check the provided data")
	}
	if err != nil {
		return []models.Credencial{}, err
	}

	return credentials, nil
}

// UpdateCredential updates a credential's information in the database
func (repo *Credential) UpdateCredential(credentialID uint64, credential models.Credencial) error {
	resultado, err := repo.db.Exec(
		`UPDATE Credenciais SET Descricao=?, SiteUrl=?, Login=?, Senha=? WHERE id=?`,
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
		return errors.New("no rows affected, check the provided data")
	}

	return nil
}

// DeleteCredential delete a credential from the database
func (repo *Credential) DeleteCredential(credentialID uint64) error {
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
		return errors.New("no rows affected, check the provided data")
	}

	return nil
}
