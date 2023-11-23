package logsCatalogados

import "errors"

var (
	//Logs
	LogsUsuario_UsuarioExistente    = "o usuário já existe"
	LogsUsuario_UsuarioNaoExistente = "usuário não encontrado"

	//Erros de consulta
	ErroUsuario_GenericoConsulta      = errors.New("ocorreu um erro ao buscar o usuário")
	ErroUsuario_UsuarioNaoCadastradao = errors.New("ocorreu um erro ao cadastrar o usuário, usuario não encontrado mesmo após o cadastro. Entre em contato com o suporte")

	//Erros ao preparar o usuário
	ErroUsuario_PrepararCadastro = errors.New("ocorreu um erro ao preparar o usuário para cadastro")
	ErroUsuario_PrepararConsulta = errors.New("ocorreu um erro ao preparar o usuário para consulta")

	//Erros de cadastro
	ErroUsuario_Cadastro = errors.New("ocorreu um erro ao realizar o cadastro do usuário")

	ErroUsuario_JsonInvalido = errors.New("ocorreu um erro ao captar os dados do usuário, verifique a estrutura e dados")

	//Dados incorretos ou vazios
	ErroUsuario_NomeVazio     = errors.New("o nome do usuário é obrigatório")
	ErroUsuario_EmailVazio    = errors.New("o email do usuário é obrigatório")
	ErroUsuario_SenhaVazia    = errors.New("a senha do usuário é obrigatório")
	ErroUsuario_EmailInvalido = errors.New("o formato de email do usuário inválido")
)
