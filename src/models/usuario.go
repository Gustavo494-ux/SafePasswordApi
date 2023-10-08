package models

import (
	"errors"
	"safePasswordApi/src/configs"
	enum "safePasswordApi/src/enum/geral"
	"safePasswordApi/src/security/encrypt/asymmetrical"
	hashEncrpt "safePasswordApi/src/security/encrypt/hash"
	symmetrical "safePasswordApi/src/security/encrypt/symmetrical"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID         uint64    `json:"id,omitempty" db:"id"`
	Nome       string    `json:"nome,omitempty" db:"nome"`
	Email      string    `json:"email,omitempty" db:"email"`
	Email_Hash string    `json:"email_hash,omitempty" db:"email_hash"`
	Senha      string    `json:"senha,omitempty" db:"senha,omitempty"`
	CriadoEm   time.Time `json:"criadoEm,omitempty" db:"criadoEm"`
}

// Prepare chamará métodos para validar e formatar o usuário recebido com base no tipo
func (usuario *Usuario) Preparar(TipoPreparacao enum.TipoPreparacao) error {
	switch TipoPreparacao {
	case enum.TipoPreparacao_Cadastro:
		if err := usuario.Validar(enum.TipoValidacao(TipoPreparacao)); err != nil {
			return err
		}
	}

	if err := usuario.Formatar(enum.TipoFormatacao(TipoPreparacao)); err != nil {
		return err
	}

	return nil
}

// Validar verifica se os campos do usuário são válidos com base no estágio determinado.
func (usuario *Usuario) Validar(TipoValidacao enum.TipoValidacao) error {
	if usuario.Nome == "" {
		return errors.New("nome é obrigatório e não pode ficar em branco")
	}

	if usuario.Email == "" {
		return errors.New("email é obrigatório e não pode ficar em branco")
	}

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("formato de email inválido")
	}

	if usuario.Senha == "" && (TipoValidacao == enum.TipoValidacao_Cadastro || TipoValidacao == enum.TipoValidacao_Atualizar) {
		return errors.New("senha é obrigatório e não pode ficar em branco")
	}

	return nil
}

// Formatar aplica a formatação necessária para cada tipo
func (usuario *Usuario) Formatar(TipoFormatacao enum.TipoFormatacao) error {
	var err error
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)

	switch TipoFormatacao {
	case enum.TipoFormatacao_Cadastro:
	case enum.TipoFormatacao_Atualizar:
		{
			usuario.Senha, err = hashEncrpt.GenerateSHA512(usuario.Senha)
			if err != nil {
				return err
			}

			usuario.Email_Hash, err = hashEncrpt.GenerateSHA512(usuario.Email)
			if err != nil {
				return err
			}

			if err = usuario.Criptografar(); err != nil {
				return err
			}
		}
	case enum.TipoFormatacao_Consulta:
		if err := usuario.Descriptografar(); err != nil {
			return err
		}
	}

	return nil
}

// CriptografarAES criptografa os dados do usuário usando AES.
func (usuario *Usuario) CriptografarAES() error {
	var err error
	if usuario.Nome, err = symmetrical.EncryptDataAES(usuario.Nome, configs.AESKey); err != nil {
		return err
	}

	if usuario.Email, err = symmetrical.EncryptDataAES(usuario.Email, configs.AESKey); err != nil {
		return err
	}

	return nil
}

// DescriptografarAES descriptografa os dados do usuário usando AES.
func (usuario *Usuario) DescriptografarAES() error {
	var err error
	if usuario.Nome, err = symmetrical.DecryptDataAES(usuario.Nome, configs.AESKey); err != nil {
		return err
	}

	if usuario.Email, err = symmetrical.DecryptDataAES(usuario.Email, configs.AESKey); err != nil {
		return err
	}

	return nil
}

// CriptografarRSA criptografa os dados do usuário usando RSA.
func (usuario *Usuario) CriptografarRSA() error {
	var err error
	publicKey, err := asymmetrical.ParseRSAPublicKey(configs.RSAPublicKey)
	if err != nil {
		return err
	}

	if usuario.Nome, err = asymmetrical.EncryptRSA(usuario.Nome, publicKey); err != nil {
		return err
	}

	if usuario.Email, err = asymmetrical.EncryptRSA(usuario.Email, publicKey); err != nil {
		return err
	}
	return nil
}

// DescriptografarRSA descriptografa os dados do usuário usando RSA.
func (usuario *Usuario) DescriptografarRSA() error {
	var err error
	privateKey, err := asymmetrical.ParseRSAPrivateKey(configs.RSAPrivateKey)
	if err != nil {
		return err
	}

	if usuario.Nome, err = asymmetrical.DecryptRSA(usuario.Nome, privateKey); err != nil {
		return err
	}

	if usuario.Email, err = asymmetrical.DecryptRSA(usuario.Email, privateKey); err != nil {
		return err
	}
	return nil
}

// Descriptografar descriptografa os dados do usuário usando criptografia AES e RSA.
func (usuario *Usuario) Descriptografar() error {
	err := usuario.DescriptografarRSA()
	if err != nil {
		return err
	}

	err = usuario.DescriptografarAES()
	if err != nil {
		return err
	}

	return nil
}

// Criptografar criptografa os dados do usuário usando criptografia RSA e AES
func (usuario *Usuario) Criptografar() error {
	err := usuario.CriptografarAES()
	if err != nil {
		return err
	}

	err = usuario.CriptografarRSA()
	if err != nil {
		return err
	}

	return nil
}
