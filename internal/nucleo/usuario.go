package nucleo

type Usuario struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo" gorm:"unique"`
	Clave  string `json:"clave"`
}
