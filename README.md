# 🌐 Orbe Users API

API RESTful en Golang para la gestión de usuarios. Desarrollada como parte de una prueba técnica para el perfil de Desarrollador Backend Jr. con enfoque en arquitectura limpia y creatividad en la estructura del proyecto.

## 🚀 Tecnologías usadas

- [Golang](https://go.dev/) – Lenguaje principal
- [Gin Gonic](https://gin-gonic.com/) – Framework HTTP
- [GORM](https://gorm.io/) – ORM para Go
- [PostgreSQL](https://www.postgresql.org/) – Base de datos relacional
- [Thunder Client](https://www.thunderclient.com/) – Cliente REST en VS Code

## 📁 Estructura del proyecto

```text
orbe-users/
├── cmd/                  # Punto de entrada (main.go)
├── internal/
│   ├── conexion/         # Conexión a la base de datos
│   ├── nucleo/           # Modelos de dominio (usuario)
│   ├── archivo/          # Repositorio de datos
│   ├── sentido/          # Servicios / lógica de negocio
│   └── pasarela/         # Handlers HTTP (controladores)
├── docs/                 # Documentación de API
├── migrations/           # Script SQL para crear tablas
├── go.mod / go.sum       # Dependencias
└── README.md             # Documentación principal
```

## 📦 Instalación y ejecución

### 1. Clonar el repositorio

```bash
git clone https://github.com/filosocode/orbe-users.git
cd orbe-users
```

### 2. Configurar PostgreSQL

Asegúrate de tener PostgreSQL instalado y ejecutándose en el puerto `5432`.

Crear la base de datos `orbe`:

```sql
CREATE DATABASE orbe;
```

### 3. Ejecutar el proyecto

```bash
go run ./cmd/orbe.go
```

> GORM realizará la migración automática y creará la tabla `usuarios`.

## 📌 Endpoints disponibles

| Método | Ruta             | Descripción                |
|--------|------------------|----------------------------|
| GET    | `/`              | Ruta raíz de prueba        |
| POST   | `/usuarios`      | Crear un nuevo usuario     |
| GET    | `/usuarios`      | Obtener lista de usuarios  |
| GET    | `/usuarios/:id`  | Obtener usuario por ID     |
| PUT    | `/usuarios/:id`  | Actualizar usuario por ID  |
| DELETE | `/usuarios/:id`  | Eliminar usuario por ID    |

## 🧪 Ejemplo con Thunder Client

### 🟢 Crear usuario (`POST /usuarios`)

**Headers:**
- `Content-Type: application/json`

**Body:**

```json
{
  "nombre": "Andrés",
  "correo": "andres@brm.com",
  "clave": "1234"
}
```

### 🔎 Obtener todos los usuarios (`GET /usuarios`)

**URL:**

```
http://localhost:8080/usuarios
```

**Respuesta esperada:**

```json
[
  {
    "id": 1,
    "nombre": "Andrés",
    "correo": "andres@brm.com",
    "clave": "1234"
  }
]
```

### 🔎 Obtener usuario por ID (`GET /usuarios/1`)

**URL:**

```
http://localhost:8080/usuarios/1
```

### 🛠 Actualizar usuario (`PUT /usuarios/1`)

**Headers:**
- `Content-Type: application/json`

**Body:**

```json
{
  "nombre": "Andrés Muñoz",
  "correo": "andres@brm.com",
  "clave": "secreta"
}
```

### 🗑 Eliminar usuario (`DELETE /usuarios/1`)

**URL:**

```
http://localhost:8080/usuarios/1
```

## 📘 ¿Qué aprendí?

- Cómo estructurar una API en Go usando Clean Architecture
- Integrar Gin con GORM y PostgreSQL
- Separar lógica por capas: controlador, servicio y repositorio
- Migraciones automáticas y uso de GORM
- Pruebas de endpoints REST usando Thunder Client

## 🔧 Qué mejoraría

- Encriptar contraseñas antes de almacenarlas
- Validaciones más completas en los handlers
- Agregar autenticación JWT
- Crear tests automatizados con `testing` y `testify`
- Agregar Docker y Docker Compose para facilitar el despliegue

## 🧠 Autor

**Andrés Muñoz**  
GitHub: [@filosocode](https://github.com/filosocode)

## 🏁 Licencia

Este proyecto fue desarrollado exclusivamente para fines evaluativos y de aprendizaje.
