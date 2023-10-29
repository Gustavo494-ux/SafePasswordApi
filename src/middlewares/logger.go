package middlewares

import (
	"fmt"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/security/auth"

	"github.com/labstack/echo/v4"
)

type requisicao struct {
	Url         string
	Ip          string
	Token       string
	TokenValido bool
	UsuarioId   int
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.Logger().Info("Requisição realizada", montarRequisicao(c))
		return nil
	}
}

func montarRequisicao(c echo.Context) requisicao {
	token := auth.ExtrairToken(c)
	var valido bool
	if auth.ValidarToken(c) == nil {
		valido = true
	}
	idUsuario, err := auth.ExtrairUsuarioID(c)
	if err != nil && valido {
		logger.Logger().Error(fmt.Sprintf("Ocorreu um erro ao tentar extrair o UsuarioId do token %s", token), err)
		return requisicao{}
	}

	r := requisicao{
		Url:         c.Request().Header.Get("curl"),
		Ip:          c.RealIP(),
		Token:       token,
		TokenValido: valido,
		UsuarioId:   int(idUsuario),
	}
	return r
}
