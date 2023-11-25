package repository

import (
	"safePasswordApi/src/core/domain/usuario"
	"safePasswordApi/src/core/errors"
)

type IUsuarioRepository interface {
	BuscarTodosUsuarios() ([]usuario.Usuario, errors.Error)
}
