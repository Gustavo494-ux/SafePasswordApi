package log

import (
	"flag"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/modules/GerenciadordeJson"
	"safePasswordApi/src/utility/fileHandler"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type NivelLog zerolog.Level

const (
	NivelLog_Debug        NivelLog = NivelLog(zerolog.DebugLevel)
	NivelLog_Desabilitado NivelLog = NivelLog(zerolog.Disabled)
	NivelLog_Informacoes  NivelLog = NivelLog(zerolog.InfoLevel)
	NivelLog_Erro         NivelLog = NivelLog(zerolog.ErrorLevel)

	NivelLog_Panico       NivelLog = NivelLog(zerolog.PanicLevel)
	NivelLog_Rastreamento NivelLog = NivelLog(zerolog.TraceLevel)
)

func ConfigurarLog(nivelLog NivelLog) {
	zerolog.TimeFieldFormat = configs.FormatoDataHora
	log.Logger = log.With().Caller().Logger()

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if nivelLog < -1 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func init() {
	ConfigurarLog(NivelLog_Rastreamento)
	arquivoLog, err := fileHandler.CarregarArquivo(configs.CaminhoArquivoLog)
	if err != nil {
		Fatal("Erro ao carregar arquivo de log", err)
	}
	logger := zerolog.New(arquivoLog).With().Timestamp().Logger()
	log.Logger = logger
}

func Info(mensagem string, dados ...interface{}) {
	log.Info().
		Caller(1).
		Str("Dados Adicionais", converterSliceDadosParaJsonString(dados)).
		Msg(mensagem)
}

func Error(mensagem string, err error, dados ...interface{}) {
	log.Error().
		Caller(1).
		Err(err).
		Str("Dados Adicionais", converterSliceDadosParaJsonString(dados)).
		Msg(mensagem)
}

func Fatal(mensagem string, err error, dados ...interface{}) {
	log.Fatal().
		Caller(1).
		Err(err).
		Str("Dados Adicionais", converterSliceDadosParaJsonString(dados)).
		Msg(mensagem)
}

func converterSliceDadosParaJsonString(dados ...interface{}) (jsonString string) {
	var dado string
	var erroLocal error
	for _, Valor := range dados {
		dado, erroLocal = GerenciadordeJson.InterfaceParaJsonString(Valor)
		if erroLocal != nil {
			Error("Ocorreu um erro ao converter uma interface para json string", erroLocal, dado)
			return
		}
		jsonString += dado
	}
	return
}
