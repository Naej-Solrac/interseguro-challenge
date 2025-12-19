package main

import (
	"challenge-go/internal/adapters/handler"
	"challenge-go/internal/adapters/nodeclient"
	"challenge-go/internal/core/service"
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// 0. Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se encontró archivo .env, usando variables del sistema")
	}

	// 1. Inicializar Fiber
	app := fiber.New()

	// 2. Configuración desde variables de entorno
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("❌ JWT_SECRET no está configurado")
	}

	nodeServiceURL := os.Getenv("NODE_SERVICE_URL")
	if nodeServiceURL == "" {
		nodeServiceURL = "http://localhost:4000" // valor por defecto
	}

	adminUser := os.Getenv("ADMIN_USERNAME")
	adminPass := os.Getenv("ADMIN_PASSWORD")
	if adminUser == "" || adminPass == "" {
		log.Fatal("❌ ADMIN_USERNAME y ADMIN_PASSWORD deben estar configurados")
	}

	// 3. INYECCIÓN DE DEPENDENCIAS (Conectar los cables)
	// Adaptador Node -> Servicio -> Handler
	nodeAdapter := nodeclient.NewNodeAdapter(nodeServiceURL)
	myService := service.NewService(nodeAdapter)
	myHandler := handler.NewMatrixHandler(myService)
	loginHandler := handler.NewLoginHandler(jwtSecret, adminUser, adminPass)

	// 4. RUTAS PÚBLICAS (Login)
	app.Post("/login", loginHandler)

	// 5. MIDDLEWARE JWT (El guardia de seguridad)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(401).JSON(fiber.Map{"error": "Token inválido o expirado"})
		},
	}))

	// 6. RUTAS PRIVADAS (Procesar Matriz)
	app.Post("/process", myHandler.Process)

	// 7. Arrancar
	log.Println("🚀 Go API corriendo en el puerto 8080")
	log.Fatal(app.Listen(":8080"))
}
