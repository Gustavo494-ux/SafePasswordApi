package models

import (
	"errors"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/security/encrypt/asymmetrical"
	symmetricEncryp "safePasswordApi/src/security/encrypt/symmetrical"
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
func (credencial *Credencial) Preparar(etapa, chave string) error {
	if erro := credencial.Validar(); erro != nil {
		return erro
	}

	if erro := credencial.Formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (credencial *Credencial) Validar() error {
	if credencial.UsuarioId == 0 {
		return errors.New("usuário é obrigatório e não pode estar em branco")
	}

	if credencial.Senha == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (credencial *Credencial) Formatar(etapa string) error {
	var erro error
	switch etapa {
	case "salvarDados":
		{
			if erro = credencial.Criptografar(); erro != nil {
				return erro
			}
		}

	case "consultarDados":
		{
			if erro = credencial.Descriptografar(); erro != nil {
				return erro
			}
		}

	}
	return nil
}

func (credencial *Credencial) Criptografar() error {
	err := credencial.CriptografarAES()
	if err != nil {
		return err
	}
	err = credencial.CriptografarRSA()
	if err != nil {
		return err
	}

	return nil
}

func (credencial *Credencial) Descriptografar() error {
	err := credencial.DescriptografarRSA()
	if err != nil {
		return err
	}

	err = credencial.DescriptografarAES()
	if err != nil {
		return err
	}

	return nil
}

func (credencial *Credencial) CriptografarAES() error {
	var erro error
	if credencial.Descricao, erro = symmetricEncryp.EncryptDataAES(credencial.Descricao, configs.AESKey); erro != nil {
		return erro
	}

	if credencial.SiteUrl, erro = symmetricEncryp.EncryptDataAES(credencial.SiteUrl, configs.AESKey); erro != nil {
		return erro
	}

	if credencial.Login, erro = symmetricEncryp.EncryptDataAES(credencial.Login, configs.AESKey); erro != nil {
		return erro
	}

	if credencial.Senha, erro = symmetricEncryp.EncryptDataAES(credencial.Senha, configs.AESKey); erro != nil {
		return erro
	}

	return nil
}

func (credencial *Credencial) DescriptografarAES() error {
	var erro error
	if credencial.Descricao, erro = symmetricEncryp.DecryptDataAES(credencial.Descricao, configs.AESKey); erro != nil {
		return erro
	}

	if credencial.SiteUrl, erro = symmetricEncryp.DecryptDataAES(credencial.SiteUrl, configs.AESKey); erro != nil {
		return erro
	}

	if credencial.Login, erro = symmetricEncryp.DecryptDataAES(credencial.Login, configs.AESKey); erro != nil {
		return erro
	}

	if credencial.Senha, erro = symmetricEncryp.DecryptDataAES(credencial.Senha, configs.AESKey); erro != nil {
		return erro
	}

	return nil
}

func (credencial *Credencial) CriptografarRSA() error {
	var erro error
	publicKey, erro := asymmetrical.ParseRSAPublicKey(configs.RSAPublicKey)
	if erro != nil {
		return erro
	}

	if credencial.Descricao, erro = asymmetrical.EncryptRSA(credencial.Descricao, publicKey); erro != nil {
		return erro
	}

	if credencial.SiteUrl, erro = asymmetrical.EncryptRSA(credencial.SiteUrl, publicKey); erro != nil {
		return erro
	}

	if credencial.Login, erro = asymmetrical.EncryptRSA(credencial.Login, publicKey); erro != nil {
		return erro
	}

	if credencial.Senha, erro = asymmetrical.EncryptRSA(credencial.Senha, publicKey); erro != nil {
		return erro
	}

	return nil
}

func (credencial *Credencial) DescriptografarRSA() error {
	var erro error
	privateKey, erro := asymmetrical.ParseRSAPrivateKey(configs.RSAPrivateKey)
	if erro != nil {
		return erro
	}

	if credencial.Descricao, erro = asymmetrical.DecryptRSA(credencial.Descricao, privateKey); erro != nil {
		return erro
	}

	if credencial.SiteUrl, erro = asymmetrical.DecryptRSA(credencial.SiteUrl, privateKey); erro != nil {
		return erro
	}

	if credencial.Login, erro = asymmetrical.DecryptRSA(credencial.Login, privateKey); erro != nil {
		return erro
	}

	if credencial.Senha, erro = asymmetrical.DecryptRSA(credencial.Senha, privateKey); erro != nil {
		return erro
	}

	return nil
}
