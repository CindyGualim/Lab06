![image](https://github.com/user-attachments/assets/2247fa98-ff9f-404f-97ea-306eb8878485)

# LaLigaTracker API

Este proyecto expone un conjunto de endpoints para administrar partidos de fútbol de LaLiga.

## Endpoints

- **GET /api/matches**  
- **GET /api/matches/{id}**  
- **POST /api/matches**  
- **PUT /api/matches/{id}**  
- **DELETE /api/matches/{id}**  
- **PATCH /api/matches/{id}/goal**  
- **PATCH /api/matches/{id}/yellowcard**  
- **PATCH /api/matches/{id}/redcard**  
- **PATCH /api/matches/{id}/extratime**  

Revisa la documentación en el archivo [`swagger.yaml`](./swagger.yaml).

## Uso

1. **Ejecuta el Backend**  
   ```bash
   go run main.go
