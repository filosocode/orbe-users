# Documentación de la API - Orbe Users

API RESTful para la gestión de usuarios. A continuación se detallan todos los endpoints disponibles, junto con sus métodos, rutas, estructuras de request/response y códigos HTTP esperados.

---

## GET /usuarios

**Descripción:** Obtener la lista completa de usuarios registrados.

- **URL:** `/usuarios`
- **Método:** GET
- **Respuesta exitosa:**
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
- **Código HTTP:** `200 OK`

---

## GET /usuarios/:id

**Descripción:** Obtener un usuario específico por su ID.

- **URL:** `/usuarios/1`
- **Método:** GET
- **Respuesta exitosa:**
```json
{
  "id": 1,
  "nombre": "Andrés",
  "correo": "andres@brm.com",
  "clave": "1234"
}
```
- **Código HTTP:** `200 OK`
- **Código de error:** `404 Not Found` si el ID no existe.

---

## POST /usuarios

**Descripción:** Crear un nuevo usuario.

- **URL:** `/usuarios`
- **Método:** POST
- **Headers:**
  - Content-Type: application/json
- **Cuerpo del request:**
```json
{
  "nombre": "Andrés",
  "correo": "andres@brm.com",
  "clave": "1234"
}
```
- **Respuesta exitosa:**
```json
{
  "id": 2,
  "nombre": "Andrés",
  "correo": "andres@brm.com",
  "clave": "1234"
}
```
- **Código HTTP:** `201 Created`
- **Errores posibles:**
  - `400 Bad Request` si los campos son inválidos.
  - `500 Internal Server Error` si falla la creación.

---

## 🔹 PUT /usuarios/:id

**Descripción:** Actualizar un usuario existente.

- **URL:** `/usuarios/1`
- **Método:** PUT
- **Headers:**
  - Content-Type: application/json
- **Cuerpo del request:**
```json
{
  "nombre": "Andrés Muñoz",
  "correo": "andresmu@example.com",
  "clave": "nueva123"
}
```
- **Respuesta exitosa:**
```json
{
  "id": 1,
  "nombre": "Andrés Muñoz",
  "correo": "andresmu@example.com",
  "clave": "nueva123"
}
```
- **Código HTTP:** `200 OK`
- **Errores posibles:** `400`, `404`, `500`

---

## DELETE /usuarios/:id

**Descripción:** Eliminar un usuario por su ID.

- **URL:** `/usuarios/1`
- **Método:** DELETE
- **Respuesta exitosa:** sin cuerpo
- **Código HTTP:** `204 No Content`
- **Errores posibles:** `400`, `404`, `500`

---

## Notas generales

- Todas las respuestas son en formato JSON.
- El servidor corre en `http://localhost:8080`
- Recomendado probar con [Thunder Client](https://www.thunderclient.com/)

---

## Autor

Andrés Muñoz – [@filosocode](https://github.com/filosocode)
