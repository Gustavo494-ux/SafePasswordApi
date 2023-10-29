package inicializacao

//Inicializar: realiza as configuracões necessarias para o funcionamento do projeto
func Inicializar() {
	CarregarDotEnv()
	InicializarEncriptacao()
	InicializarMysql()
	InicializaAPI()
}
