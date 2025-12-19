// Infraestructura: Controller para estadísticas de matrices
import { Request, Response } from 'express';
import { getMatricesStatisticsUseCase } from '../application/getMatrixStats.usecase';

/**
 * Controlador para POST /stats
 */
export async function matrixStatsController(req: Request, res: Response): Promise<void> {
  try {
    const { Q, R } = req.body;
    if (!Array.isArray(Q) || !Array.isArray(R)) {
      res.status(400).json({ error: 'Q y R deben ser matrices (array de arrays).' });
      return;
    }
    const result = getMatricesStatisticsUseCase(Q, R);
    res.status(200).json(result);
  } catch (err: any) {
    if (err.message && err.message.includes('matriz')) {
      res.status(400).json({ error: err.message });
      return;
    }
    res.status(500).json({ error: 'Error interno del servidor.' });
  }
}
