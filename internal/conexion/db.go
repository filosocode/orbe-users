package conexion

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"orbe-users/internal/nucleo"
)

var DB *gorm.DB

func ConectarDB() {
	dsn := "host=db user=postgres password=BRM dbname=orbe port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	DB = db
	log.Println("✅ Conexión exitosa a PostgreSQL")

	// Auto migración del modelo Usuario
	err = db.AutoMigrate(&nucleo.Usuario{})
	if err != nil {
		log.Fatal("Error al migrar el modelo Usuario: ", err)
	}
}
