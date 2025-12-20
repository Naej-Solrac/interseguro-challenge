package handler

import (
	"challenge-go/internal/adapters/handler/validation"
	"challenge-go/internal/core/domain"
	"challenge-go/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

// MatrixHandler maneja las peticiones HTTP
type MatrixHandler struct {
	service   ports.MatrixService // Inyectamos el puerto de entrada
	validator *validation.Validator
}

// NewMatrixHandler constructor
func NewMatrixHandler(s ports.MatrixService) *MatrixHandler {
	return &MatrixHandler{
		service:   s,
		validator: validation.NewValidator(),
	}
}

// Process es la función que Fiber ejecutará
// Process godoc
// @Summary      Procesar matriz con QR
// @Description  Recibe una matriz, realiza factorización QR usando el algoritmo de Gram-Schmidt, y retorna las matrices Q y R junto con sus estadísticas (promedio, min, max, suma, es diagonal)
// @Tags         Matrix Processing
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        matrix body validation.MatrixRequest true "Matriz de entrada (array bidimensional de números float64)"
// @Success      200 {object} map[string]interface{} "Factorización QR completada con estadísticas"
// @Failure      400 {object} map[string]string "Formato JSON inválido o error de validación"
// @Failure      401 {object} map[string]string "Token JWT inválido, expirado o no proporcionado"
// @Failure      500 {object} map[string]string "Error interno al procesar la matriz"
// @Router       /matrix/process [post]
func (h *MatrixHandler) Process(c *fiber.Ctx) error {
	var req validation.MatrixRequest

	// 1. Parsear el Body (JSON)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	// 2. Validar estructura con go-playground/validator
	if err := h.validator.ValidateStruct(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// 3. Convertir al formato del dominio
	input := domain.MatrixInput{Data: req.Data}

	// 4. Llamar al Servicio (Core)
	response, err := h.service.Process(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// 5. Responder
	return c.JSON(response)
}
