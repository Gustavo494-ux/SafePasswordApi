package logger

import (
	"safePasswordApi/src/configs"
	"safePasswordApi/src/modules/GerenciadordeJson"
	"safePasswordApi/src/utility/fileHandler"

	"github.com/rs/zerolog"
)

type NivelLog zerolog.Level

type LoggerType struct {
	log zerolog.Logger
}

const (
	NivelLog_Debug        NivelLog = NivelLog(zerolog.DebugLevel)
	NivelLog_Desabilitado NivelLog = NivelLog(zerolog.Disabled)
	NivelLog_Informacoes  NivelLog = NivelLog(zerolog.InfoLevel)
	NivelLog_Erro         NivelLog = NivelLog(zerolog.ErrorLevel)

	NivelLog_Panico       NivelLog = NivelLog(zerolog.PanicLevel)
	NivelLog_Rastreamento NivelLog = NivelLog(zerolog.TraceLevel)
)

// Logger cria uma instância de logger
func Logger() *LoggerType {
	var logger LoggerType
	logger.init()
	return &logger
}

// init: realiza a configuração necessária para o pacote funcionar
func (logger *LoggerType) init() {
	logger.configurarLog(NivelLog_Rastreamento)
}

// configurarLog: realiza a configuração básica para o log funcionar
func (logger *LoggerType) configurarLog(nivelLog NivelLog) {
	zerolog.TimeFieldFormat = configs.FormatoDataHora

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if nivelLog < -1 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	arquivoLog, err := fileHandler.CarregarArquivo(configs.CaminhoArquivoLog)
	if err != nil {
		logger.Fatal("Erro ao carregar arquivo de log", err)
	}

	logger.log = zerolog.New(arquivoLog).With().Timestamp().Logger()
}

// Info: cria um log de informação
func (logger *LoggerType) Info(mensagem string, dados ...interface{}) {
	logger.log.
		Info().
		Caller(1).
		Str("Dados Adicionais", logger.converterSliceDadosParaJsonString(dados)).
		Msg(mensagem)
}

// Error: cria um log de erro
func (logger *LoggerType) Error(mensagem string, err error, dados ...interface{}) {
	logger.log.
		Error().
		Caller(1).
		Err(err).
		Str("Dados Adicionais", logger.converterSliceDadosParaJsonString(dados)).
		Msg(mensagem)
}

// Fatal: cria um log de erro fatal
func (logger *LoggerType) Fatal(mensagem string, err error, dados ...interface{}) {
	logger.log.
		Fatal().
		Caller(1).
		Err(err).
		Str("Dados Adicionais", logger.converterSliceDadosParaJsonString(dados)).
		Msg(mensagem)
}

// converterSliceDadosParaJsonString: converte uma interface para jsonString
func (logger *LoggerType) converterSliceDadosParaJsonString(dados ...interface{}) (jsonString string) {
	var dado string
	var erroLocal error
	for _, Valor := range dados {
		dado, erroLocal = GerenciadordeJson.InterfaceParaJsonString(Valor)
		if erroLocal != nil {
			logger.Error("Ocorreu um erro ao converter uma interface para json string", erroLocal)
			return
		}
		jsonString += dado
	}
	return
}
