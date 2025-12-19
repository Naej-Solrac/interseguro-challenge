package handler

import (
	"challenge-go/internal/core/domain"
	"challenge-go/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

// MatrixHandler maneja las peticiones HTTP
type MatrixHandler struct {
	service ports.MatrixService // Inyectamos el puerto de entrada
}

// NewMatrixHandler constructor
func NewMatrixHandler(s ports.MatrixService) *MatrixHandler {
	return &MatrixHandler{service: s}
}

// Process es la función que Fiber ejecutará
func (h *MatrixHandler) Process(c *fiber.Ctx) error {
	var req domain.MatrixInput
	
	// 1. Parsear el Body (JSON)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	// 2. Llamar al Servicio (Core)
	response, err := h.service.Process(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// 3. Responder
	return c.JSON(response)
}