package usuario

import (
	"time"
)

type Usuario struct {
	id         uint64
	nome       string
	email      string
	email_Hash string
	senha      string
	criadoEm   time.Time
}

func (instance *Usuario) ID() uint64 {
	return instance.id
}

func (instance *Usuario) Nome() string {
	return instance.nome
}

func (instance *Usuario) Email() string {
	return instance.email
}

func (instance *Usuario) Email_Hash() string {
	return instance.email_Hash
}

func (instance *Usuario) Senha() string {
	return instance.senha
}

func (instance *Usuario) CriadoEm() time.Time {
	return instance.criadoEm
}

func (instance *Usuario) IsZero() bool {
	return instance == &Usuario{}
}
