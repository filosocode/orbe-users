package archivo

import (
	"orbe-users/internal/conexion"
	"orbe-users/internal/nucleo"
)

// CrearUsuario inserta un nuevo usuario en la base de datos
func CrearUsuario(usuario *nucleo.Usuario) error {
	result := conexion.DB.Create(usuario)
	return result.Error
}

// ObtenerUsuarios devuelve todos los usuarios
func ObtenerUsuarios() ([]nucleo.Usuario, error) {
	var usuarios []nucleo.Usuario
	result := conexion.DB.Find(&usuarios)
	return usuarios, result.Error
}

// ObtenerUsuarioPorID busca un usuario por su ID
func ObtenerUsuarioPorID(id uint) (*nucleo.Usuario, error) {
	var usuario nucleo.Usuario
	result := conexion.DB.First(&usuario, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &usuario, nil
}

// ActualizarUsuario guarda los cambios de un usuario existente
func ActualizarUsuario(usuario *nucleo.Usuario) error {
	result := conexion.DB.Save(usuario)
	return result.Error
}

// EliminarUsuario elimina un usuario por ID
func EliminarUsuario(id uint) error {
	result := conexion.DB.Delete(&nucleo.Usuario{}, id)
	return result.Error
}
