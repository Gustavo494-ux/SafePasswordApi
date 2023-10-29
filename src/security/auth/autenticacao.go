package auth

import (
	"errors"
	"fmt"
	"safePasswordApi/src/configs"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarTokenJWT(usuarioID uint64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	permissoes := token.Claims.(jwt.MapClaims)
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	return token.SignedString([]byte(configs.SecretKeyJWT))
}

// ValidarToken verifica se o token passado na requisição é válido
func ValidarToken(c echo.Context) error {
	tokenString := ExtrairToken(c)
	_, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}
	return nil
}

// ExtrairUsuarioID retorna o usuarioId que está salvo no token
func ExtrairUsuarioID(c echo.Context) (uint64, error) {
	tokenString := ExtrairToken(c)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("token inválido")
}

// ExtrairToken: Extrai o Token da requisição
func ExtrairToken(c echo.Context) string {
	token := c.Request().Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return configs.SecretKeyJWT, nil
}
