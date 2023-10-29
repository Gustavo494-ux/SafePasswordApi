package configuracoes

import (
	"os"
	"safePasswordApi/src/modules/tipo"
	"strconv"
)

type API struct {
	Porta *int
}

// ConfigurarApi: realiza as configurações necessárias para o funcionametno da API
func (api *API) ConfigurarApi() {
	api.carregarVariaveis()
}

// carregarVariaveis: carrega as variaveis necessarias
func (api *API) carregarVariaveis() {
	*api.Porta, _ = strconv.Atoi(os.Getenv("API_PORT"))
	*api.Porta = tipo.Coalesce().Int(*api.Porta, 5000)
}
