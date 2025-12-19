// Infraestructura: Definición de rutas REST
import { Router } from 'express';
import { matrixStatsController } from './matrixStats.controller';

const router = Router();

router.post('/stats', matrixStatsController);

export default router;
