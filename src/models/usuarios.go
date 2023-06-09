package models

import (
	"errors"
	"fmt"
	"safePasswordApi/src/security"
	"strconv"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty" db:"id"`
	Nome     string    `json:"nome,omitempty" db:"nome"`
	Email    string    `json:"email,omitempty" db:"email"`
	Senha    string    `json:"senha,omitempty" db:"senha"`
	CriadoEm time.Time `json:"criadoEm,omitempty" db:"criadoem"`
}

// Preparar vai chamar os métodos para validar e formatar usuário  recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar("cadastro"); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o email inserido é inválido")
	}

	if usuario.Senha == "" && etapa == "cadastro" {
		return errors.New("o eenha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaHash, erro := security.GerarHash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = senhaHash
		fmt.Println(usuario.Senha)
	}

	return nil
}

func (usuario *Usuario) GerarChaveDeCodificacaoSimetrica() ([]byte, error) {
	idHash, erro := security.GerarHash(strconv.FormatUint(usuario.ID, 10))
	if erro != nil {
		return []byte{}, erro
	}

	var senhaHash string
	if len(usuario.Senha) == 128 {
		senhaHash = usuario.Senha
	} else {
		senhaHash, erro = security.GerarHash(usuario.Senha)
		if erro != nil {
			return []byte{}, erro
		}
	}

	chaveDeCodificacao, erro := security.GerarHash(fmt.Sprintf(idHash, usuario.ID, senhaHash))
	if erro != nil {
		return []byte{}, erro
	}
	fmt.Println(chaveDeCodificacao)
	return []byte(chaveDeCodificacao), nil
}
