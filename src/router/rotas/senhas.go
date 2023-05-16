package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasSenhas = []Rota{
	{
		URI:                    "/senhas",
		Method:                 http.MethodPost,
		Function:               controllers.CriarUsuario,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/senhas/senhaId",
		Method:                 http.MethodGet,
		Function:               controllers.CriarUsuario,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/senhas/usuarioId",
		Method:                 http.MethodGet,
		Function:               controllers.CriarUsuario,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/senhas/senhaId",
		Method:                 http.MethodPut,
		Function:               controllers.CriarUsuario,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/senhas/senhaId",
		Method:                 http.MethodDelete,
		Function:               controllers.CriarUsuario,
		RequiresAuthentication: true,
	},
}
