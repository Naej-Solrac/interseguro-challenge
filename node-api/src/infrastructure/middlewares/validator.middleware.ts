// Middleware genérico para validar con Joi
import { Request, Response, NextFunction } from 'express';
import { Schema } from 'joi';

/**
 * Middleware genérico de validación con Joi
 * @param schema - Esquema de Joi específico para validar
 * @returns Middleware de Express
 */
export const validate = (schema: Schema) => {
  return (req: Request, res: Response, next: NextFunction): void => {
    const { error } = schema.validate(req.body, { abortEarly: false });

    if (error) {
      const errors = error.details.map((detail) => detail.message);
      res.status(400).json({
        error: 'Validación fallida',
        details: errors,
      });
      return;
    }

    next();
  };
};
