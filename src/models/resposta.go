package models

import (
	"fmt"
	"safePasswordApi/src/modules/logger"

	"github.com/labstack/echo/v4"
)

type Resposta struct {
	StatusCode            int
	Mensagem              string
	Error                 error
	DadosAdicionaisLogger []interface{}
	Contexto              echo.Context
}

// RespostaRequisicao cria uma instância resposta
func RespostaRequisicao(c echo.Context) *Resposta {
	return &Resposta{Contexto: c}
}

// Erro: configura a resposta para um erro
func (resposta *Resposta) Erro(StatusCode int, erro error, MensagemLogger string, DadosAdicionaisLogger ...interface{}) *Resposta {
	resposta.StatusCode = StatusCode
	resposta.Error = erro
	resposta.Mensagem = fmt.Sprintf("%s Id Requisição: %s", MensagemLogger, resposta.Contexto)
	resposta.DadosAdicionaisLogger = DadosAdicionaisLogger
	logger.Logger().Error(resposta.Mensagem, resposta.Error, resposta.DadosAdicionaisLogger...)
	return resposta
}

func (resposta *Resposta) JSON() error {
	return resposta.Contexto.JSON(resposta.StatusCode, resposta.Error.Error())
}

func (resposta *Resposta) String() error {
	return resposta.Contexto.JSON(resposta.StatusCode, resposta.Error.Error())
}
