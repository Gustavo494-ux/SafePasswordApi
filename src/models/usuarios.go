package models

import (
	"errors"
	"fmt"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/security/encrypt/asymmetrical"
	hashEncrpt "safePasswordApi/src/security/encrypt/hash"
	symmetricEncrypt "safePasswordApi/src/security/encrypt/symmetrical"
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
	if erro := usuario.Validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.Formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) Validar(etapa string) error {
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

func (usuario *Usuario) Formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)

	switch etapa {
	case "cadastro":
		{
			senhaHash, err := hashEncrpt.GenerateSHA512(usuario.Senha)
			if err != nil {
				return err
			}
			usuario.Senha = senhaHash
			// if err = usuario.Encrypt(); err != nil {
			// 	return err
			// }
		}
		// case "consulta":
		// 	if err := usuario.Decrypt(); err != nil {
		// 		return err
		// 	}
	}

	return nil
}

func (usuario *Usuario) GerarChaveDeCodificacaoSimetrica() ([]byte, error) {
	idHash, erro := hashEncrpt.GenerateSHA512(strconv.FormatUint(usuario.ID, 10))
	if erro != nil {
		return []byte{}, erro
	}

	var senhaHash string
	if len(usuario.Senha) == 128 {
		senhaHash = usuario.Senha
	} else {
		senhaHash, erro = hashEncrpt.GenerateSHA512(usuario.Senha)
		if erro != nil {
			return []byte{}, erro
		}
	}

	chaveDeCodificacao, erro := hashEncrpt.GenerateSHA512(fmt.Sprintf(idHash, usuario.ID, senhaHash))
	if erro != nil {
		return []byte{}, erro
	}
	return []byte(chaveDeCodificacao), nil
}

func (usuario *Usuario) EncryptAES() error {
	var err error
	if usuario.Nome, err = symmetricEncrypt.EncryptDataAES(usuario.Nome, configs.AESKey); err != nil {
		return err
	}

	if usuario.Email, err = symmetricEncrypt.EncryptDataAES(usuario.Email, configs.AESKey); err != nil {
		return err
	}

	return nil
}

func (usuario *Usuario) DecryptAES() error {
	var err error
	if usuario.Nome, err = symmetricEncrypt.DecryptDataAES(usuario.Nome, configs.AESKey); err != nil {
		return err
	}

	if usuario.Email, err = symmetricEncrypt.DecryptDataAES(usuario.Email, configs.AESKey); err != nil {
		return err
	}

	return nil
}

func (usuario *Usuario) EncryptRSA() error {
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

func (usuario *Usuario) DecryptRSA() error {
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

func (usuario *Usuario) Encrypt() error {
	err := usuario.EncryptAES()
	if err != nil {
		return err
	}

	err = usuario.EncryptRSA()
	if err != nil {
		return err
	}

	return nil
}

func (usuario *Usuario) Decrypt() error {
	err := usuario.DecryptRSA()
	if err != nil {
		return err
	}

	err = usuario.DecryptAES()
	if err != nil {
		return err
	}

	return nil
}
