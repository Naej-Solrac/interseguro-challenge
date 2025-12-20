// Esquema de validación específico para el endpoint /stats
import Joi from 'joi';

/**
 * Esquema de validación para matrices Q y R
 * Valida que:
 * - Q y R sean arrays
 * - Cada elemento sea un array de números
 * - Las matrices no estén vacías
 */
export const matrixStatsSchema = Joi.object({
  Q: Joi.array()
    .items(
      Joi.array()
        .items(Joi.number().required())
        .min(1)
        .required()
    )
    .min(1)
    .required()
    .messages({
      'array.base': 'Q debe ser una matriz (array de arrays)',
      'array.min': 'Q no puede estar vacía',
      'any.required': 'Q es obligatoria',
    }),

  R: Joi.array()
    .items(
      Joi.array()
        .items(Joi.number().required())
        .min(1)
        .required()
    )
    .min(1)
    .required()
    .messages({
      'array.base': 'R debe ser una matriz (array de arrays)',
      'array.min': 'R no puede estar vacía',
      'any.required': 'R es obligatoria',
    }),
});
