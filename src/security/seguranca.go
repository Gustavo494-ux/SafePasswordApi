package security

import "golang.org/x/crypto/bcrypt"

// GerarHash uma string e colocar um hash nela
func GerarHash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarHash compara uma string e um hash e retorna se elas são iguais
func VerificarSenha(senhaHash string, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
