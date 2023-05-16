package models

import (
	"errors"
	"time"
)

type Senha struct {
	Id        int64     `json:"id,omitempty" db:"id"`
	UsuarioId int64     `json:"usuarioId,omitempty" db:"usuarioid"`
	Nome      string    `json:"nome,omitempty" db:"Nome"`
	Senha     string    `json:"senha,omitempty" db:"senha"`
	CriadoEm  time.Time `json:"criadoEm,omitempty" db:"criadoEm,omitempty"`
}

// Validar realiza a verificação se cada campo está devidamente preenchido
func (senha *Senha) Validar() error {
	if senha.UsuarioId == 0 {
		return errors.New("usuário o qual a senha pertence não foi informado")
	}
	if senha.Nome == "" {
		return errors.New("nome da senha é obrigatorio")
	}
	if senha.Senha == "" {
		return errors.New("senha é obrigatoria")
	}
	return nil
}
