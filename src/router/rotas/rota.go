package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotasLogin)

	for _, rota := range rotas {
		if rota.RequiresAuthentication {
			r.HandleFunc(rota.URI, middlewares.Autenticar(rota.Function)).Methods(rota.Method)
		} else {
			r.HandleFunc(rota.URI, rota.Function).Methods(rota.Method)
		}
	}
	return r
}
