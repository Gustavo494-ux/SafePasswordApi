package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                    "/usuarios",
		Method:                 http.MethodPost,
		Function:               controllers.CriarUsuario,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios",
		Method:                 http.MethodGet,
		Function:               controllers.BuscarUsuarios,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodGet,
		Function:               controllers.BuscarUsuario,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodPut,
		Function:               controllers.AtualizarUsuario,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletarUsuario,
		RequiresAuthentication: true,
	},
}
