// Dominio: Lógica pura para estadísticas de matrices


import { MatrixStatistics } from './dto/MatrixStatistics.dto';

/**
 * Calcula estadísticas de una matriz numérica.
 * @param matrix Matriz de números.
 * @returns Estadísticas de la matriz.
 */
export function calculateMatrixStatistics(matrix: number[][]): MatrixStatistics {
  if (!Array.isArray(matrix) || matrix.length === 0 || !Array.isArray(matrix[0])) {
    throw new Error('La matriz debe ser un array de arrays no vacío.');
  }
  const rowCount = matrix.length;
  const colCount = matrix[0].length;
  let max = -Infinity;
  let min = Infinity;
  let sum = 0;
  let count = 0;
  let isDiagonal = true;
  for (let i = 0; i < rowCount; i++) {
    if (!Array.isArray(matrix[i]) || matrix[i].length !== colCount) {
      throw new Error('Todas las filas deben tener la misma longitud.');
    }
    for (let j = 0; j < colCount; j++) {
      const value = matrix[i][j];
      if (typeof value !== 'number' || isNaN(value)) {
        throw new Error('Todos los elementos deben ser números.');
      }
      max = Math.max(max, value);
      min = Math.min(min, value);
      sum += value;
      count++;
      if (i === j) {
        if (value === 0) isDiagonal = false;
      } else {
        if (value !== 0) isDiagonal = false;
      }
    }
  }
  const avg = sum / count;
  return { max, min, avg, sum, isDiagonal };
}
