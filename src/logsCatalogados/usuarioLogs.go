package logsCatalogados

import "errors"

var (
	ErroUsuario_PrepararCadastro = errors.New("ocorreu um erro ao preparar o usuário para cadastro")
	ErroUsuario_NomeVazio        = errors.New("o nome do usuário é obrigatório")
	ErroUsuario_EmailVazio       = errors.New("o email do usuário é obrigatório")
	ErroUsuario_SenhaVazia       = errors.New("a senha do usuário é obrigatório")
	ErroUsuario_EmailInvalido    = errors.New("o formato de email do usuário inválido")
)
