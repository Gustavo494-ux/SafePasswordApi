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
		RequiresAuthentication: false,
	},
	{
		URI:                    "/usuarios",
		Method:                 http.MethodGet,
		Function:               controllers.BuscarUsuarios,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodGet,
		Function:               controllers.BuscarUsuario,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodPut,
		Function:               controllers.AtualizarUsuario,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletarUsuario,
		RequiresAuthentication: false,
	},
}
