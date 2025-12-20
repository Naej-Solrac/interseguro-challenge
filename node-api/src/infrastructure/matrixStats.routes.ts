// Infraestructura: Definición de rutas REST
import { Router } from 'express';
import { matrixStatsController } from './matrixStats.controller';
import { validate } from './middlewares/validator.middleware';
import { matrixStatsSchema } from './schemas/matrixStats.schema';

const router = Router();

// Aplicamos el middleware de validación con el esquema específico
router.post('/stats', validate(matrixStatsSchema), matrixStatsController);

export default router;
