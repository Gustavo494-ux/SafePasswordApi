package inicializacao

import (
	"safePasswordApi/src/configs"
	configuracoes "safePasswordApi/src/routines/Configuracoes"
)

// InicializarMysql: realiza as configurações necessarias para o Mysql
func InicializarMysql() {
	Mysql := configuracoes.BancoDeDados{
		StringConexao: &configs.StringConnection,
	}
	Mysql.ConfigurarBancoDeDados()
}
