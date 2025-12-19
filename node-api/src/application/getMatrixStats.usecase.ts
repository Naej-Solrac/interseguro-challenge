// Aplicación: Caso de uso para obtener estadísticas de dos matrices
import { calculateMatrixStatistics } from '../domain/stats.logic';
import { MatrixStatistics } from '../domain/dto/MatrixStatistics.dto';

export interface MatricesStatisticsResult {
  Q: MatrixStatistics;
  R: MatrixStatistics;
}

/**
 * Procesa dos matrices y retorna sus estadísticas.
 * @param Q Primera matriz.
 * @param R Segunda matriz.
 * @returns Estadísticas de ambas matrices.
 */
export function getMatricesStatisticsUseCase(Q: number[][], R: number[][]): MatricesStatisticsResult {
  return {
    Q: calculateMatrixStatistics(Q),
    R: calculateMatrixStatistics(R)
  };
}
