package inicializacao

import (
	"safePasswordApi/src/configs"
	configuracoes "safePasswordApi/src/routines/Configuracoes"
)

// InicializaAPI: realiza as configurações necessarias para a API
func InicializaAPI() {
	api := configuracoes.API{
		Porta: &configs.Port,
	}
	api.ConfigurarApi()
}
