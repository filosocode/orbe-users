package main

import (
	"log"

	"orbe-users/internal/conexion"
	"orbe-users/internal/pasarela"

	"github.com/gin-gonic/gin"
)

func main() {
	conexion.ConectarDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"mensaje": "Bienvenido a orbe-users 🚀"})
	})

	pasarela.RegistrarRutasUsuario(r)

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	r.Run()
}
