package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Validator es un wrapper del go-playground/validator
type Validator struct {
	validate *validator.Validate
}

// NewValidator crea una nueva instancia del validador
func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// ValidateStruct valida un struct y retorna un error de Fiber si falla
func (v *Validator) ValidateStruct(data interface{}) error {
	if err := v.validate.Struct(data); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf(
				"'%s' %s",
				err.Field(),
				getValidationMessage(err),
			))
		}
		return fiber.NewError(400, strings.Join(errors, ", "))
	}
	return nil
}

// getValidationMessage retorna un mensaje legible según el tag de validación
func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "es obligatorio"
	case "email":
		return "debe ser un email válido"
	case "min":
		return fmt.Sprintf("debe tener mínimo %s caracteres", err.Param())
	case "max":
		return fmt.Sprintf("debe tener máximo %s caracteres", err.Param())
	case "gte":
		return fmt.Sprintf("debe ser mayor o igual a %s", err.Param())
	case "lte":
		return fmt.Sprintf("debe ser menor o igual a %s", err.Param())
	default:
		return fmt.Sprintf("falló la validación '%s'", err.Tag())
	}
}
