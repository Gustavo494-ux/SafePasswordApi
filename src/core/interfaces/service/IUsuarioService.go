package service

import (
	"safePasswordApi/src/core/domain/usuario"
	"safePasswordApi/src/core/errors"
)

type IUsuarioService interface {
	BuscarTodosUsuarios() ([]usuario.Usuario, errors.Error)
}
