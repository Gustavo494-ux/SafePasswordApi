package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasLogin = Rota{
	URI:                    "/Login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	RequiresAuthentication: false,
}
