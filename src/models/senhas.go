package models

import (
	"errors"
	"time"
)

type Senha struct {
	Id        uint      `json:"id,omitempty" db:"id"`
	UsuarioId uint      `json:"usuarioId,omitempty" db:"usuarioId"`
	Nome      string    `json:"Nome,omitempty" db:"Nome"`
	Senha     string    `json:"senha,omitempty" db:"senha"`
	CriadoEm  time.Time `json:"criadoEm,omitempty" db:"criadoEm"`
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
