![image](https://github.com/user-attachments/assets/2247fa98-ff9f-404f-97ea-306eb8878485)
![image](https://github.com/user-attachments/assets/90c9e64b-2842-463f-b0fe-bc43e7a978d0)
# LaLigaTracker Backend

Este proyecto implementa un API REST para gestionar partidos de LaLiga. Incluye los siguientes endpoints:

- **GET /api/matches:** Obtener todos los partidos.
- **GET /api/matches/{id}:** Obtener un partido por ID.
- **POST /api/matches:** Crear un nuevo partido.
- **PUT /api/matches/{id}:** Actualizar un partido existente.
- **DELETE /api/matches/{id}:** Eliminar un partido.
- **PATCH /api/matches/{id}/goals:** Incrementar goles.
- **PATCH /api/matches/{id}/yellowcards:** Registrar tarjeta amarilla.
- **PATCH /api/matches/{id}/redcards:** Registrar tarjeta roja.
- **PATCH /api/matches/{id}/extratime:** Registrar tiempo extra.

## Documentación

- [Swagger Documentation](./swagger.yaml)

## Uso de la API

Revisa el archivo `llms.txt` para más detalles sobre el uso de la API.

## Colección Postman / Hoppscotch

He creado una colección para probar el API. Puedes importarla en Postman o acceder a ella en Hoppscotch a través del siguiente enlace:

[Link a Colección Postman/Hoppscotch](https://hoppscotch.io/)

## Configuración CORS

El backend está configurado para permitir solicitudes CORS desde cualquier origen.

## Cómo Ejecutar el Proyecto

1. Asegúrate de tener Go instalado.
2. Clona o descarga el repositorio.
3. Abre la terminal en la carpeta del proyecto y ejecuta:
   ```bash
   go run main.go

