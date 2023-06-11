package repository

import (
	"errors"
	"fmt"
	"safePasswordApi/src/models"

	"github.com/jmoiron/sqlx"
)

type Credencial struct {
	db *sqlx.DB
}

// NovoRepositoDeCredencial cria um repositório de Credencial
func NovoRepositoDeCredencial(db *sqlx.DB) *Credencial {
	return &Credencial{db}
}

// CriarCredencial Adiciona um nova credencial
func (repositorio Credencial) CriarCredencial(credencial models.Credencial) (uint64, error) {
	fmt.Println(credencial)
	statement, erro := repositorio.db.Exec(
		` INSERT INTO Credenciais (UsuarioId,Descricao,siteUrl,Login,Senha ) values (?,?,?,?,?) `,
		credencial.UsuarioId,
		credencial.Descricao,
		credencial.SiteUrl,
		credencial.Login,
		credencial.Senha,
	)
	linhasAfetadas, err := statement.RowsAffected()
	if linhasAfetadas == 0 {
		fmt.Println(err)
		return 0, errors.New("nenhuma linha foi afetada, verifique os dados passados")
	}
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	if erro != nil {
		fmt.Println(err)
		return 0, erro
	}

	credencialID, erro := statement.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(credencialID), nil
}

// BuscarCredencialPorId busca uma credencial pelo ID
func (repositorio Credencial) BuscarCredencialPorId(credencialId uint64) (models.Credencial, error) {
	credencial := models.Credencial{}
	erro := repositorio.db.Get(&credencial,
		` SELECT UsuarioId,Descricao,siteUrl,Login,Senha,CriadoEm from Credenciais WHERE id = ? `,
		credencialId,
	)

	if credencial.Id == 0 {
		return models.Credencial{}, errors.New("nenhuma credencial foi encontrada, verifique os dados passados")
	}

	if erro != nil {
		return models.Credencial{}, erro
	}
	return credencial, nil
}

// BuscarCredenciais busca todas as credenciais salvas no banco
func (repositorio Credencial) BuscarCredenciais() ([]models.Credencial, error) {
	var credenciais []models.Credencial
	erro := repositorio.db.Select(&credenciais, "SELECT UsuarioId,Descricao,siteUrl,Login,Senha,CriadoEm from Credenciais FROM Credenciais ")
	if len(credenciais) == 0 {
		return []models.Credencial{}, errors.New("nenhuma credencial foi encontrada, verifique os dados fornecidos")
	}

	if erro != nil {
		return []models.Credencial{}, erro
	}
	return credenciais, nil
}

// AtualizarCredencial Atualiza as informações de uma credencial no banco
func (repositorio Credencial) AtualizarCredencial(credencialId uint64, credencial models.Credencial) error {
	statement, erro := repositorio.db.Exec(
		` UPDATE Credenciais SET Descricao =?, SiteUrl =?, Login =?,Senha = ? WHERE id =? `,
		credencial.Descricao,
		credencial.SiteUrl,
		credencial.Login,
		credencial.Senha,
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

// DeletarCredencial deleta uma credencial do banco de dados
func (repositorio Credencial) DeletarCredencial(credencialId uint64) error {
	statement, erro := repositorio.db.Exec(
		` DELETE FROM Credenciais WHERE id =? `,
		credencialId,
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
