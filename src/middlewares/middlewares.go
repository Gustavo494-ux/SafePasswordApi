package middlewares

import (
	"api/src/authentication"
	"api/src/respostas"
	"net/http"
)

func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w, r)
	}
}
