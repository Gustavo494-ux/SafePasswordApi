package models

import (
	"errors"
	"safePasswordApi/src/configs"
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

func (credencial *Credencial) Prepare(step string) error {
	if err := credencial.Validate(); err != nil {
		return err
	}

	if err := credencial.Format(step); err != nil {
		return err
	}

	return nil
}

func (credencial *Credencial) Validate() error {
	if credencial.UsuarioId == 0 {
		return errors.New("user is required and cannot be blank")
	}

	if credencial.Senha == "" {
		return errors.New("password is required and cannot be blank")
	}

	return nil
}

func (credencial *Credencial) Format(step string) error {
	var err error
	switch step {
	case "saveData":
		if err = credencial.Encrypt(); err != nil {
			return err
		}

	case "retrieveData":
		if err = credencial.Decrypt(); err != nil {
			return err
		}
	}

	return nil
}

func (credencial *Credencial) Encrypt() error {
	err := credencial.EncryptAES()
	if err != nil {
		return err
	}

	err = credencial.EncryptRSA()
	if err != nil {
		return err
	}

	return nil
}

func (credencial *Credencial) Decrypt() error {
	err := credencial.DecryptRSA()
	if err != nil {
		return err
	}

	err = credencial.DecryptAES()
	if err != nil {
		return err
	}

	return nil
}

func (credencial *Credencial) EncryptAES() error {
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

func (credencial *Credencial) DecryptAES() error {
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

func (credencial *Credencial) EncryptRSA() error {
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

func (credencial *Credencial) DecryptRSA() error {
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
