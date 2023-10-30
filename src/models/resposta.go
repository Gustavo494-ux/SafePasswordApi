package models

import (
	"fmt"
	"safePasswordApi/src/modules/logger"

	"github.com/labstack/echo/v4"
)

type Resposta struct {
	StatusCode            int
	Mensagem              string
	DadosAdicionaisLogger []interface{}
	Contexto              echo.Context
	dadoRetorno           interface{}
}

// RespostaRequisicao cria uma instância resposta
func RespostaRequisicao(c echo.Context) *Resposta {
	return &Resposta{Contexto: c}
}

// Erro: configura a resposta para resposta para uma requisição bem sucedida
func (resposta *Resposta) Sucesso(StatusCode int, DadosResposta interface{}, MensagemLogger string, DadosAdicionaisLogger ...interface{}) *Resposta {
	resposta.StatusCode = StatusCode
	resposta.dadoRetorno = DadosResposta
	resposta.Mensagem = fmt.Sprintf("%s , Id Requisição: %s, StatusCode: %d", MensagemLogger, resposta.Contexto, StatusCode)
	resposta.DadosAdicionaisLogger = DadosAdicionaisLogger
	logger.Logger().Info(resposta.Mensagem, resposta.DadosAdicionaisLogger...)
	return resposta
}

// Erro: configura a resposta para um erro
func (resposta *Resposta) Erro(StatusCode int, erro error, MensagemLogger string, DadosAdicionaisLogger ...interface{}) *Resposta {
	resposta.StatusCode = StatusCode
	resposta.dadoRetorno = erro
	resposta.Mensagem = fmt.Sprintf("%s , Id Requisição: %s, StatusCode: %d", MensagemLogger, resposta.Contexto, StatusCode)
	resposta.DadosAdicionaisLogger = DadosAdicionaisLogger
	logger.Logger().Error(resposta.Mensagem, erro, resposta.DadosAdicionaisLogger...)
	return resposta
}

// JSON: envia uma resposta para a requisição em formato JSON
func (resposta *Resposta) JSON() error {
	return resposta.Contexto.JSON(resposta.StatusCode, resposta.dadoRetorno)
}

// String: envia uma resposta para a requisição em formato String
func (resposta *Resposta) String() error {
	return resposta.Contexto.JSON(resposta.StatusCode, resposta.dadoRetorno)
}
