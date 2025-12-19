package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// NewLoginHandler crea un handler de login con la configuración inyectada
// Esto hace el código más testeable y sigue principios SOLID
func NewLoginHandler(jwtSecret, adminUser, adminPass string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type LoginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
		}

		// Validación con credenciales (en producción: BD con bcrypt)

		if req.Username == adminUser && req.Password == adminPass {
			// Crear token JWT
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": req.Username,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			})

			tokenString, err := token.SignedString([]byte(jwtSecret))
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Error al generar token"})
			}

			return c.JSON(fiber.Map{
				"token": tokenString,
			})
		}

		return c.Status(401).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}
}
