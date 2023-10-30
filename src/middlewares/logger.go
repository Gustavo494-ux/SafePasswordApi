package middlewares

import (
	"fmt"
	"safePasswordApi/src/modules/logger"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type requisicao struct {
	Url         string
	Ip          string
	Token       string
	TokenValido bool
	UsuarioId   int
}

func LoggerZeroLogPersonalizado(c echo.Context, v middleware.RequestLoggerValues) error {
	requestID := v.RequestID

	// Capturar informações do cliente
	clientIP := c.RealIP()
	userAgent := c.Request().UserAgent()

	// Capturar informações da solicitação
	httpMethod := c.Request().Method
	requestURI := c.Request().RequestURI
	requestHeaders := c.Request().Header

	// Capturar informações adicionais
	currentTime := time.Now()
	// Realizar ações na solicitação e medir o tempo de resposta
	start := time.Now()
	latency := time.Since(start)
	responseSize := c.Response().Size

	responseString := fmt.Sprintf("ID da Requisição: %s,"+
		"Informações do Cliente{"+
		"Endereço IP: %s,"+
		"User-Agent: %s,"+
		"Informações da Solicitação: "+
		"Método HTTP: %s,"+
		"URI: %s,"+
		"Cabeçalhos da Solicitação: %v,"+
		"Tempo de Resposta: %s,"+
		"Latência: %s,"+
		"Tamanho da Resposta: %d bytes,"+
		"Status do Usuário: %s "+
		" }", requestID, clientIP, userAgent, httpMethod, requestURI, requestHeaders, time.Since(start), latency, responseSize, currentTime)
	logger.Logger().Info("Requisição realizada", responseString)
	return nil
}
