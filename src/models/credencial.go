package models

import (
	"errors"
	"safePasswordApi/src/security"
	"time"
)

type Credencial struct {
	Id        uint64    `json:"id,omitempty" db:"id"`
	UsuarioId uint64    `json:"usuarioId,omitempty" db:"usuarioId"`
	Descricao string    `json:"descricao,omitempty" db:"descricao"`
	SiteUrl   string    `json:"siteUrl,omitempty" db:"siteUrl"`
	Login     string    `json:"login,omitempty" db:"login"`
	Senha     string    `json:"senha,omitempty" db:"senha"`
	CriadoEm  time.Time `json:"criadoEm,omitempty" db:"criadoem"`
}

// Preparar vai chamar os métodos para validar e formatar credencial  recebido
func (credencial *Credencial) Preparar(etapa string, chave string) error {
	if erro := credencial.validar(); erro != nil {
		return erro
	}

	if erro := credencial.formatar(etapa, chave); erro != nil {
		return erro
	}
	return nil
}

func (credencial *Credencial) validar() error {
	if credencial.UsuarioId == 0 {
		return errors.New("usuário é obrigatório e não pode estar em branco")
	}

	if credencial.Senha == "" {
		return errors.New("o senha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (credencial *Credencial) formatar(etapa, chave string) error {
	var erro error
	switch etapa {
	case "salvarDados":
		{
			if erro = credencial.criptografar(chave); erro != nil {
				return erro
			}
		}

	case "consultarDados":
		{
			if erro = credencial.descriptografar(chave); erro != nil {
				return erro
			}
		}

	}
	return nil
}

func (credencial *Credencial) criptografar(chave string) error {
	var erro error
	if credencial.Descricao, erro = security.CriptografarTexto(credencial.Descricao, chave); erro != nil {
		return erro
	}

	if credencial.SiteUrl, erro = security.CriptografarTexto(credencial.SiteUrl, chave); erro != nil {
		return erro
	}

	if credencial.Login, erro = security.CriptografarTexto(credencial.Login, chave); erro != nil {
		return erro
	}

	if credencial.Senha, erro = security.CriptografarTexto(credencial.Senha, chave); erro != nil {
		return erro
	}

	return nil
}

func (credencial *Credencial) descriptografar(chave string) error {
	var erro error
	if credencial.Descricao, erro = security.DescriptografarTexto(credencial.Descricao, chave); erro != nil {
		return erro
	}

	if credencial.SiteUrl, erro = security.DescriptografarTexto(credencial.SiteUrl, chave); erro != nil {
		return erro
	}

	if credencial.Login, erro = security.DescriptografarTexto(credencial.Login, chave); erro != nil {
		return erro
	}

	if credencial.Senha, erro = security.DescriptografarTexto(credencial.Senha, chave); erro != nil {
		return erro
	}

	return nil
}
