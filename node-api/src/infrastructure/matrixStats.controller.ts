// Infraestructura: Controller para estadísticas de matrices
import { Request, Response } from 'express';
import { getMatricesStatisticsUseCase } from '../application/getMatrixStats.usecase';

/**
 * @openapi
 * /api/stats:
 *   post:
 *     summary: Calcula estadísticas de matrices Q y R
 *     description: Recibe dos matrices (Q y R) provenientes de una factorización QR y devuelve estadísticas de cada una (promedio, min, max, suma, es diagonal)
 *     tags:
 *       - Statistics
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - Q
 *               - R
 *             properties:
 *               Q:
 *                 $ref: '#/components/schemas/Matrix'
 *                 description: Matriz Q de la factorización QR
 *               R:
 *                 $ref: '#/components/schemas/Matrix'
 *                 description: Matriz R de la factorización QR
 *           example:
 *             Q: [[0.169, 0.897], [0.507, 0.276], [0.845, -0.345]]
 *             R: [[5.916, 7.437], [0, 0.828]]
 *     responses:
 *       200:
 *         description: Estadísticas calculadas exitosamente
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 Q:
 *                   $ref: '#/components/schemas/MatrixStatistics'
 *                 R:
 *                   $ref: '#/components/schemas/MatrixStatistics'
 *             example:
 *               Q:
 *                 avg: 0.392
 *                 min: -0.345
 *                 max: 0.897
 *                 sum: 2.349
 *                 isDiagonal: false
 *               R:
 *                 avg: 3.545
 *                 min: 0
 *                 max: 7.437
 *                 sum: 14.181
 *                 isDiagonal: false
 *       400:
 *         description: Error de validación o datos inválidos
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/ErrorResponse'
 *       500:
 *         description: Error interno del servidor
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 error:
 *                   type: string
 *                   example: Error interno del servidor.
 */
export async function matrixStatsController(req: Request, res: Response): Promise<void> {
  try {
    // Ya está validado por Joi, podemos confiar en los datos
    const { Q, R } = req.body;
    
    const result = getMatricesStatisticsUseCase(Q, R);
    res.status(200).json(result);
  } catch (err: any) {
    // Manejo de errores del dominio (lógica de negocio)
    if (err.message && err.message.includes('matriz')) {
      res.status(400).json({ error: err.message });
      return;
    }
    res.status(500).json({ error: 'Error interno del servidor.' });
  }
}
