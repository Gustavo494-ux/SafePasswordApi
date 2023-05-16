package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasSenhas = []Rota{
	{
		URI:                    "/senhas",
		Method:                 http.MethodPost,
		Function:               controllers.CriarSenha,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/senhas/senhaId",
		Method:                 http.MethodGet,
		Function:               controllers.BuscarSenhaPorId,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/senhas/usuarioId",
		Method:                 http.MethodGet,
		Function:               controllers.BuscarSenhaPorUsuario,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/senhas/senhaId",
		Method:                 http.MethodPut,
		Function:               controllers.AtualizarSenha,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/senhas/senhaId",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletarSenha,
		RequiresAuthentication: true,
	},
}
