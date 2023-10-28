package inicializacao

import (
	"safePasswordApi/src/configs"
	configuracoes "safePasswordApi/src/routines/Configuracoes"
)

// inicializarMysql: realiza as configurações necessarias para o Mysql
func inicializarMysql() {
	Mysql := configuracoes.BancoDeDados{
		StringConexao: &configs.StringConnection,
	}
	Mysql.ConfigurarBancoDeDados()
}
