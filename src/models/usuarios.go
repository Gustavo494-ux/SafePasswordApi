package models

import "time"

type Usuario struct {
	ID       int64     `json:"id,omitempty" db:"id"`
	Nome     string    `json:"nome,omitempty" db:"nome"`
	Nick     string    `json:"nick,omitempty" db:"nick"`
	Email    string    `json:"email,omitempty" db:"email"`
	Senha    string    `json:"senha,omitempty" db:"senha"`
	CriadoEm time.Time `json:"criadoEm,omitempty" db:"criado_em"`
}
