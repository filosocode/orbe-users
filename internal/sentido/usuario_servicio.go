package sentido

import (
	"orbe-users/internal/archivo"
	"orbe-users/internal/nucleo"
)

func CrearUsuario(usuario *nucleo.Usuario) error {
	return archivo.CrearUsuario(usuario)
}

func ObtenerUsuarios() ([]nucleo.Usuario, error) {
	return archivo.ObtenerUsuarios()
}

func ObtenerUsuarioPorID(id uint) (*nucleo.Usuario, error) {
	return archivo.ObtenerUsuarioPorID(id)
}

func ActualizarUsuario(usuario *nucleo.Usuario) error {
	return archivo.ActualizarUsuario(usuario)
}

func EliminarUsuario(id uint) error {
	return archivo.EliminarUsuario(id)
}
