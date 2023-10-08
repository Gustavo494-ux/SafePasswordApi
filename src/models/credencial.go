package models

import (
	"errors"
	"safePasswordApi/src/configs"
	enum "safePasswordApi/src/enum/geral"
	"safePasswordApi/src/security/encrypt/asymmetrical"
	symmetricEncrypt "safePasswordApi/src/security/encrypt/symmetrical"
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

// Preparar chamará métodos para validar e formatar a credencial recebido com base no tipo
func (credencial *Credencial) Preparar(TipoPreparacao enum.TipoPreparacao) error {
	if err := credencial.Validar(); err != nil {
		return err
	}

	if err := credencial.Formatar(enum.TipoFormatacao(TipoPreparacao)); err != nil {
		return err
	}

	return nil
}

// Validar verifica se os campos da credencial são válidos com base no tipo determinado.
func (credencial *Credencial) Validar() error {
	if credencial.UsuarioId == 0 {
		return errors.New("o usuário é obrigatório e não pode ficar em branco")
	}

	if credencial.Senha == "" {
		return errors.New("a senha é obrigatória e não pode ficar em branco")
	}

	return nil
}

// Formatar aplica a formatação necessária para cada tipo
func (credencial *Credencial) Formatar(TipoFormatacao enum.TipoFormatacao) error {
	var err error
	switch TipoFormatacao {
	case enum.TipoFormatacao_Cadastro,
		enum.TipoFormatacao_Atualizar:
		if err = credencial.Criptografar(); err != nil {
			return err
		}

	case enum.TipoFormatacao_Consulta:
		if err = credencial.Descriptografar(); err != nil {
			return err
		}
	}

	return nil
}

// Criptografar criptografa os dados da credencial usando criptografia AES e RSA
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

// Criptografar criptografa os dados da credencial usando criptografia RSA e AES
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

// CriptografarAES criptografa os dados da credencial usando AES.
func (credencial *Credencial) CriptografarAES() error {
	var err error
	if credencial.Descricao, err = symmetricEncrypt.EncryptDataAES(credencial.Descricao, configs.AESKey); err != nil {
		return err
	}

	if credencial.SiteUrl, err = symmetricEncrypt.EncryptDataAES(credencial.SiteUrl, configs.AESKey); err != nil {
		return err
	}

	if credencial.Login, err = symmetricEncrypt.EncryptDataAES(credencial.Login, configs.AESKey); err != nil {
		return err
	}

	if credencial.Senha, err = symmetricEncrypt.EncryptDataAES(credencial.Senha, configs.AESKey); err != nil {
		return err
	}

	return nil
}

// DescriptografarAES descriptografa os dados da credencial usando AES.
func (credencial *Credencial) DescriptografarAES() error {
	var err error
	if credencial.Descricao, err = symmetricEncrypt.DecryptDataAES(credencial.Descricao, configs.AESKey); err != nil {
		return err
	}

	if credencial.SiteUrl, err = symmetricEncrypt.DecryptDataAES(credencial.SiteUrl, configs.AESKey); err != nil {
		return err
	}

	if credencial.Login, err = symmetricEncrypt.DecryptDataAES(credencial.Login, configs.AESKey); err != nil {
		return err
	}

	if credencial.Senha, err = symmetricEncrypt.DecryptDataAES(credencial.Senha, configs.AESKey); err != nil {
		return err
	}

	return nil
}

// CriptografarRSA criptografa os dados da credencial usando RSA.
func (credencial *Credencial) CriptografarRSA() error {
	var err error
	publicKey, err := asymmetrical.ParseRSAPublicKey(configs.RSAPublicKey)
	if err != nil {
		return err
	}

	if credencial.Descricao, err = asymmetrical.EncryptRSA(credencial.Descricao, publicKey); err != nil {
		return err
	}

	if credencial.SiteUrl, err = asymmetrical.EncryptRSA(credencial.SiteUrl, publicKey); err != nil {
		return err
	}

	if credencial.Login, err = asymmetrical.EncryptRSA(credencial.Login, publicKey); err != nil {
		return err
	}

	if credencial.Senha, err = asymmetrical.EncryptRSA(credencial.Senha, publicKey); err != nil {
		return err
	}

	return nil
}

// DescriptografarRSA descriptografa os dados da credencial usando RSA.
func (credencial *Credencial) DescriptografarRSA() error {
	var err error
	privateKey, err := asymmetrical.ParseRSAPrivateKey(configs.RSAPrivateKey)
	if err != nil {
		return err
	}

	if credencial.Descricao, err = asymmetrical.DecryptRSA(credencial.Descricao, privateKey); err != nil {
		return err
	}

	if credencial.SiteUrl, err = asymmetrical.DecryptRSA(credencial.SiteUrl, privateKey); err != nil {
		return err
	}

	if credencial.Login, err = asymmetrical.DecryptRSA(credencial.Login, privateKey); err != nil {
		return err
	}

	if credencial.Senha, err = asymmetrical.DecryptRSA(credencial.Senha, privateKey); err != nil {
		return err
	}

	return nil
}
