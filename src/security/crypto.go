package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
)

// GerarHash gera um hash para a senha informada
func GerarHash(senhaDescriptografada string) (string, error) {
	h := sha512.New()
	_, err := h.Write([]byte(senhaDescriptografada))
	if err != nil {
		return "", err
	}
	senhaCriptografada := h.Sum(nil)
	hashHex := hex.EncodeToString(senhaCriptografada)
	return hashHex, nil
}

// CompararHash verifica se a senha informada corresponde ao hash
func CompararHash(senhaCriptografada, senhaDescriptografada string) error {
	hashDescriptografado, err := GerarHash(senhaDescriptografada)
	if err != nil {
		return err
	}
	senhaCriptografadaBytes := []byte(senhaCriptografada)

	if fmt.Sprintf("%x", hashDescriptografado) != fmt.Sprintf("%x", senhaCriptografadaBytes) {
		return errors.New("os dados de login estão são inválidos")
	}

	return nil
}

// CriptografarTexto Criptografar a informação
func CriptografarTexto(textoClaro string, chave string) (string, error) {
	chaveByte := reduzirChave([]byte(chave), 32)
	iv := reduzirChave([]byte(chave), aes.BlockSize)

	bloco, err := aes.NewCipher(chaveByte)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(bloco, iv)
	textoCifrado := make([]byte, len(textoClaro))
	stream.XORKeyStream(textoCifrado, []byte(textoClaro))

	textoCifrado = append(iv, textoCifrado...)
	return fmt.Sprintf("%x", textoCifrado), nil
}

// CriptografarTexto Descriptografar a informação
func DescriptografarTexto(textoCifrado string, chave string) (string, error) {
	chaveByte := reduzirChave([]byte(chave), 32)
	iv := reduzirChave([]byte(chave), aes.BlockSize)

	textoCifradoBytes, err := hex.DecodeString(textoCifrado)
	if err != nil {
		return "", err
	}

	if len(textoCifradoBytes) < aes.BlockSize {
		return "", errors.New("texto cifrado inválido")
	}

	bloco, err := aes.NewCipher(chaveByte)
	if err != nil {
		return "", err
	}

	textoCifradoBytes = textoCifradoBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(bloco, iv)
	stream.XORKeyStream(textoCifradoBytes, textoCifradoBytes)

	return string(textoCifradoBytes), nil
}

func reduzirChave(chave []byte, novoTamanho int) []byte {
	novaChave, _ := GerarHash(hex.EncodeToString(chave))
	if novoTamanho >= len(novaChave) {
		return chave
	}
	return append([]byte(nil), novaChave[:novoTamanho]...)
}
