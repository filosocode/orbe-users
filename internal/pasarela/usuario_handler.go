package pasarela

import (
	"net/http"
	"strconv"

	"orbe-users/internal/nucleo"
	"orbe-users/internal/sentido"

	"github.com/gin-gonic/gin"
)

// RegistrarRutasUsuario agrega las rutas al router
func RegistrarRutasUsuario(r *gin.Engine) {
	usuarios := r.Group("/usuarios")
	{
		usuarios.POST("/", crearUsuario)
		usuarios.GET("/", obtenerUsuarios)
		usuarios.GET("/:id", obtenerUsuarioPorID)
		usuarios.PUT("/:id", actualizarUsuario)
		usuarios.DELETE("/:id", eliminarUsuario)
	}
}

func crearUsuario(c *gin.Context) {
	var usuario nucleo.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sentido.CrearUsuario(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, usuario)
}

func obtenerUsuarios(c *gin.Context) {
	usuarios, err := sentido.ObtenerUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func obtenerUsuarioPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	usuario, err := sentido.ObtenerUsuarioPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func actualizarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var usuario nucleo.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario.ID = uint(id)

	if err := sentido.ActualizarUsuario(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func eliminarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := sentido.EliminarUsuario(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
