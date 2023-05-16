package models

import (
	"errors"
	"time"
)

type Senha struct {
	Id        int64     `json:"id,omitempty" db:"id,omitempty"`
	UsuarioId int64     `json:"usuarioId,omitempty" db:"usuarioid,omitempty"`
	Descricao string    `json:"descricao,omitempty" db:"descricao,omitempty"`
	Login     string    `json:"login,omitempty" db:"login,omitempty"`
	Senha     string    `json:"senha,omitempty" db:"senha,omitempty"`
	CriadoEm  time.Time `json:"criadoEm,omitempty" db:"criadoEm,omitempty"`
}

// Validar realiza a verificação se cada campo está devidamente preenchido
func (senha *Senha) Validar() error {
	if senha.UsuarioId == 0 {
		return errors.New("usuário o qual a senha pertence não foi informado")
	}
	if senha.Descricao == "" {
		return errors.New("nome da senha é obrigatorio")
	}
	if senha.Senha == "" {
		return errors.New("senha é obrigatoria")
	}
	return nil
}
