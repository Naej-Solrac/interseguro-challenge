# Challenge Interseguro

Este proyecto muestra una arquitectura de microservicios usando Go y Node.js. La idea es que el servidor en Go se encargue de hacer la factorización QR de matrices (usando Gram-Schmidt), mientras que el de Node calcula las estadísticas.

## Estructura General

El sistema tiene dos APIs que se comunican entre sí:

- **Go API** (puerto 8080): Recibe las matrices, las factoriza usando QR, y luego pide las estadísticas al servicio de Node

- **Node API** (puerto 3000): Calcula las estadísticas (promedio, suma, mínimo, máximo, etc.)

```
Cliente → Go API (QR + JWT) → Node API (Stats) → Respuesta completa
```

## Proyectos

### Go API (`go-api/`)
API principal construida con Fiber v2. Implementa arquitectura hexagonal para mantener la lógica de negocio separada. Usa JWT para autenticación.

**Endpoints:**
- `POST /login` - Login para obtener token
- `POST /process` - Envía una matriz, la factoriza y devuelve todo (necesitas el token)

### Node API (`node-api/`)
Servicio secundario en Express + TypeScript. Se encarga solo de calcular estadísticas de las matrices Q y R que vienen de Go.

**Endpoints:**
- `POST /api/stats` - Recibe matrices y devuelve sus estadísticas

## Instalación

### Requisitos
- Go 1.25+
- Node.js 18+
- npm o yarn

### 1. Clonar repositorio
```bash
git clone https://github.com/Naej-Solrac/interseguro-challenge.git
cd interseguro-challenge
```

### Configurar el .env

En la carpeta `go-api/` crea un archivo `.env` con tus credenciales:

```env
JWT_SECRET=tu_secreto_jwt
ADMIN_USERNAME=tu_usuario
ADMIN_PASSWORD=tu_password
NODE_SERVICE_URL=http://localhost:3000/api
```

(No subas este archivo a git, ya está en .gitignore)

### Instalar dependencias

**Node API:**
```bash
cd node-api
npm install
```

**Go API:**
```bash
cd go-api
go mod tidy
```

### Correr todo

Necesitas dos terminales:

**Terminal 1 - Node:**
```bash
cd node-api
npm run dev
```

**Terminal 2 - Go:**
```bash
cd go-api
go run cmd/api/main.go
```

Ya Cómo usar

### Primero necesitas el token

### 1. Obtener token JWT
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"TU_USUARIO","password":"TU_PASSWORD"}'
```

**Respuesta:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Procesar una matriz
```bash
curl -X POST http://localhost:8080/process \
  -H "Authorization: Bearer <tu-token>" \
  -H "Content-Type: application/json" \
  -d '{"data":[[1,2],[3,4],[5,6]]}'
```

**Respuesta:**
```json
{
  "qr_factorization": {
    "Q": [[0.169, 0.897], [0.507, 0.276], [0.845, -0.345]],
    "R": [[5.916, 7.437], [0, 0.828]]
  },
  "statistics": {
    "Q": {
      "avg": 0.392,
      "min": -0.345,
      "max": 0.897,
      "sum": 2.349,
      "isDiagonal": false
    },
    "R": {
      "avg": 3.545,
      "min": 0,
      "max": 7.437,
      "sum": 14.181,
      "isDiagonal": false
    }
  }
}
```

## Tests

Para el servicio de Node:
```bash
cd node-api
npm test
```

Para Go:
```bash
cd go-api
go test ./...
```

LosArquitectura del código

### Go API (Hexagonal

### Go API (Hexagonal/Clean Architecture)

```
go-api/
├── cmd/api/main.go              # Punto de entrada
├── internal/
│   ├── adapters/                # Capa externa
│   │   ├── handler/             # HTTP handlers
│   │   │   ├── http.go          # Handler /process
│   │   │   └── login.go         # Handler /login
│   │   └── nodeclient/          # Cliente Node.js
│   │       └── client.go
│   └── core/                    # Capa interna (lógica de negocio)
│       ├── domain/              # Modelos
│       │   └── matrix.go
│       ├── logic/               # Algoritmos puros
│       │   ├── qr.go            # Gram-Schmidt
│       │   └── qr_test.go       # Tests unitarios
│       ├── ports/               # Interfaces (contratos)
│       │   └── ports.go
│       └── service/             # Orquestación
│           └── matrix_service.go
```Capas

### Node API (Layered Architecture)

```
node-api/
├── src/
│   ├── index.ts                 # Punto de entrada
│   ├── application/             # Casos de uso
│   │   └── getMatrixStats.usecase.ts
│   ├── domain/                  # Lógica de negocio
│   │   ├── stats.logic.ts       # Algoritmos de estadísticas
│   │   ├── dto/
│   │   │   └── MatrixStatistics.dto.ts
│   │   └── __tests__/
│   │       └── stats.logic.test.ts  # Tests con Jest
│   └── infrastructure/          # Capa externa
│       ├── server.ts            # Configuración Express
│       ├── matrixStats.routes.ts
│       └── matrixStats.controller.ts
├── jest.config.js
└── tsconfig.json
```Cosas de seguridad

- Las credenciales están en el `.env` que no se sube a git
- El login devuelve un JWT que dura 24 horas
- El endpoint `/process` está protegido, necesitas el token

## Tecnologías usadmínimo privilegio

##  Tecnologías

### Go API
- `github.com/gofiber/fiber/v2` - Framework web
- `github.com/golang-jwt/jwt/v5` - JWT
- `github.com/joho/godotenv` - Variables de entorno

### Node API
- `express` - Framework web
- `typescript` - Tipado estático
- `jest` - Testing

##  Habilidades Demostradas
Lo que incluye el proyecto

- Arquitectura hexagonal en Go (más fácil de testear y mantener)
- Comunicación entre microservicios
- Autenticación con JWT
- Tests unitarios
- Algoritmo de Gram-Schmidt para QR

---

Proyecto desarrollado para el proceso de selección de Interseguro