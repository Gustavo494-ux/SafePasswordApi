package models

import (
	"errors"
	"safePasswordApi/src/security"
	"time"
)

type Credenciais struct {
	Id        uint64    `json:"id,omitempty" db:"id"`
	UsuarioId uint64    `json:"usuarioId,omitempty" db:"usuarioId"`
	Descricao string    `json:"descricao,omitempty" db:"descricao"`
	SiteUrl   string    `json:"siteUrl,omitempty" db:"siteUrl"`
	Login     string    `json:"login,omitempty" db:"login"`
	Senha     string    `json:"senha,omitempty" db:"senha"`
	CriadoEm  time.Time `json:"criadoEm,omitempty" db:"criadoem"`
}

// Preparar vai chamar os métodos para validar e formatar usuário  recebido
func (Credenciais *Credenciais) Preparar(etapa string, chave string) error {
	if erro := Credenciais.validar(); erro != nil {
		return erro
	}

	if erro := Credenciais.formatar(etapa, chave); erro != nil {
		return erro
	}
	return nil
}

func (credenciais *Credenciais) validar() error {
	if credenciais.UsuarioId == 0 {
		return errors.New("usuário é obrigatório e não pode estar em branco")
	}

	if credenciais.Senha == "" {
		return errors.New("o senha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (credenciais *Credenciais) formatar(etapa, chave string) error {
	var erro error
	switch etapa {
	case "salvarDados":
		{
			if erro = credenciais.criptografar(chave); erro != nil {
				return erro
			}
		}

	case "consultarDados":
		{
			if erro = credenciais.descriptografar(chave); erro != nil {
				return erro
			}
		}

	}
	return nil
}

func (credenciais *Credenciais) criptografar(chave string) error {
	var erro error
	if credenciais.Descricao, erro = security.CriptografarTexto(credenciais.Descricao, chave); erro != nil {
		return erro
	}

	if credenciais.SiteUrl, erro = security.CriptografarTexto(credenciais.SiteUrl, chave); erro != nil {
		return erro
	}

	if credenciais.Login, erro = security.CriptografarTexto(credenciais.Login, chave); erro != nil {
		return erro
	}

	if credenciais.Senha, erro = security.CriptografarTexto(credenciais.Senha, chave); erro != nil {
		return erro
	}

	return nil
}

func (credenciais *Credenciais) descriptografar(chave string) error {
	var erro error
	if credenciais.Descricao, erro = security.DescriptografarTexto(credenciais.Descricao, chave); erro != nil {
		return erro
	}

	if credenciais.SiteUrl, erro = security.DescriptografarTexto(credenciais.SiteUrl, chave); erro != nil {
		return erro
	}

	if credenciais.Login, erro = security.DescriptografarTexto(credenciais.Login, chave); erro != nil {
		return erro
	}

	if credenciais.Senha, erro = security.DescriptografarTexto(credenciais.Senha, chave); erro != nil {
		return erro
	}

	return nil
}
