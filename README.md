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

**📚 Documentación:** [http://localhost:8080/api-docs/index.html](http://localhost:8080/api-docs/index.html)

**Endpoints:**
- `POST /api/v1/auth/login` - Login para obtener token JWT
- `POST /api/v1/matrix/process` - Factorización QR de matrices (requiere JWT)

### Node API (`node-api/`)
Servicio secundario en Express + TypeScript. Calcula estadísticas de las matrices Q y R.

**📚 Documentación:** [http://localhost:3000/api-docs](http://localhost:3000/api-docs)

**Endpoints:**
- `POST /api/stats` - Calcula estadísticas de matrices

## 🚀 Instalación (con Docker - RECOMENDADO)

### Requisitos
- Docker
- Docker Compose

### 1. Clonar repositorio
```bash
git clone https://github.com/Naej-Solcar/interseguro-challenge.git
cd interseguro-challenge
```

### 2. Levantar todo con un solo comando
```bash
do**Go API:** `http://localhost:8080`
- **Node API:** `http://localhost:3000`

### 📚 Documentación interactiva (Swagger)
- **Go API Docs:** [http://localhost:8080/api-docs/index.html](http://localhost:8080/api-docs/index.html)
- **Node API Docs:** [http://localhost:3000/api-docs](http://localhost:3000/api-docs)

> 💡 **Tip:** Usa Swagger UI para probar todos los endpoints directamente desde el navegador

**¡Listo!** 🎉 Ambos servicios estarán corriendo:
- Go API en `http://localhost:8080`
- Node API en `http://localhost:3000`

### Credenciales por defecto
- **Usuario:** `admin`
- **Password:** `password123`

*(Puedes cambiarlas en el `docker-compose.yml`)*

---

## 📦 Instalación Manual (sin Docker)

<details>
<summary>Click aquí si prefieres correr sin Docker</summary>

### Requisitos
- Go 1.25+
- Node.js 18+

### 1. Instalar dependencias

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

### 2. Configurar .env

En `go-api/` crea un archivo `.env`:

```env
JWT_SECRET=secreto_para_firmar_tokens
ADMIN_USERNAME=admin
ADMIN_PASSWORD=password123
NODE_SERVICE_URL=http://localhost:3000/api
```

### 3. Correr ambos servicios

**Terminal 1 - Node:**
```bash
cd node-api
npm run dev
```

**Terminal 2 - Go:**
```bash
cd go-api
go run cmd/api/main.go
``` la API

### Opción 1: Swagger UI (Recomendado) 🎨

La forma más fácil es usar la documentación interactiva de Swagger:

1. **Levanta los servicios:**
   ```bash
   docker-compose up
   ```

2. **Abre Swagger en tu navegador:**
   - Go API: [http://localhost:8080/api-docs/index.html](http://localhost:8080/api-docs/index.html)
   - Node API: [http://localhost:3000/api-docs](http://localhost:3000/api-docs)

3. **Prueba los endpoints directamente desde la interfaz** - Swagger te permite ejecutar requests sin usar curl

### Opción 2: Ejemplo rápido con curl

**1. Login:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password123"}'
```

**2. Procesar matriz (copia el token del paso anterior):**
```bash
curl -X POST http://localhost:8080/api/v1/matrix/process \
  -H "Authorization: Bearer <tu-token>" \
  -H "Content-Type: application/json" \
  -d '{"data":[[1,2],[3,4],[5,6]]}'
```

> 📖 Para ver todos los ejemplos de request/response, consulta la documentación de Swagger   "max": 7.437,
      "sum": 14.181,
      "isDiagonal": false
    }
  }
}
```

## 🧪 Tests

**Con Docker:**
```bash
# Node API
docker-compose exec node-api npm test

# Go API
docker-compose exec go-api go test ./...
```

**Sin Docker:**
```bash
# Node API
cd node-api && npm test

# Go API
cd go-api && go test ./...
```

## 🏗️ Arquitectura del código

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
```

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
```

## 🔒 Seguridad

- JWT con expiración de 24 horas
- Endpoint `/process` protegido con token
- Las credenciales se configuran en variables de entorno
- Arquitectura de microservicios con red privada en Docker

## 🛠️ Tecnologías

### Go API
- `github.com/gofiber/fiber/v2` - Framework web
- `github.com/golang-jwt/jwt/v5` - JWT
- `github.com/joho/godotenv` - Variables de entorno

### Node API
- `express` - Framework web
- `typescript` - Tipado estático
- `jest` - Testing
# DevOps
- `Docker` & `Docker Compose` - Containerización
- Multi-stage builds para optimizar imágenes

## ✨ Características

- ✅ Arquitectura hexagonal en Go
- ✅ Comunicación entre microservicios
- ✅ Autenticacgo-playground/validator/v10` - Validación de datos
- `github.com/swaggo/fiber-swagger` - Documentación OpenAPI

### Node API
- `express` - Framework web
- `typescript` - Tipado estático
- `joi` - Validación de esquemas
- `swagger-jsdoc` & `swagger-ui-express` - Documentación OpenAPI
- `jest` - Testing

##
Con Docker:
```bash(Go) y por capas (Node)
- ✅ Comunicación entre microservicios
- ✅ Autenticación JWT con middleware
- ✅ Validación de datos (go-playground/validator + Joi)
- ✅ Documentación OpenAPI/Swagger interactiva
- ✅ Tests unitarios con cobertura
- ✅ Dockerizado y listo para producción
- ✅ Versionado de API (v1)
- ✅ Algoritmo de Gram-Schmidt para factorización
Proyecto desarrollado para el proceso de selección de Interseguro 🚀
Proyecto desarrollado para el proceso de selección de Interseguro